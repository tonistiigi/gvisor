package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	if err := unbazel(); err != nil {
		fmt.Printf("error: %+v", err)
		os.Exit(1)
	}
}

func unbazel() error {
	if _, err := os.Stat("WORKSPACE"); err != nil {
		return errors.Errorf("failed to find WORKSPACE needs to bu run in gvisor source directory")
	}

	if os.Getenv("GIT_WORK_TREE") == "" {
		if out, err := exec.Command("git", "status", "--porcelain").Output(); err != nil || len(out) != 0 {
			return errors.Errorf("not a clean git state for manual run: %s", out)
		}
	}

	if err := os.Mkdir("extractgopath", 0711); err != nil && !os.IsExist(err) {
		return err
	}

	if err := ioutil.WriteFile("extractgopath/BUILD", []byte(`load("@io_bazel_rules_go//go:def.bzl", "go_path")

go_path(
	name = "extract",
	mode = "copy",
	deps = ["//runsc"],
)
`), 0600); err != nil {
		return err
	}

	cmd := exec.Command("bazel", "build", "extractgopath:extract")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = blankEnv()
	if err := cmd.Run(); err != nil {
		return err
	}

	deps, err := readWorkspaceDeps()
	if err != nil {
		return err
	}

	deps2, err := readProtoDeps()
	if err != nil {
		return err
	}

	deps = append(deps, deps2...)

	const out = "./gvisor-out"

	os.RemoveAll(out)

	cp := exec.Command("cp", "-a", "bazel-bin/extractgopath/extract/src/gvisor.googlesource.com/gvisor", out)
	cp.Stdout = os.Stdout
	cp.Stderr = os.Stderr
	if err := cp.Run(); err != nil {
		return err
	}

	if err := chmod(out); err != nil {
		return err
	}

	f, err := os.Create(out + "/go.mod")
	if err != nil {
		return err
	}

	fmt.Fprint(f, `module gvisor.googlesource.com/gvisor

require (
`)

	for _, d := range deps {
		fmt.Fprintf(f, "\t%s %s\n", d.Name, d.SHA)
	}

	fmt.Fprint(f, ")\n")

	if err := f.Close(); err != nil {
		return err
	}

	if err := fixNeedsAssembly(out); err != nil {
		return err
	}

	cp = exec.Command("sh", "-c", "cp ./runsc/*.go "+out+"/runsc/")
	cp.Stdout = os.Stdout
	cp.Stderr = os.Stderr
	if err := cp.Run(); err != nil {
		return err
	}

	tidy := exec.Command("go", "mod", "tidy")
	tidy.Stdout = os.Stdout
	tidy.Stderr = os.Stderr
	tidy.Dir = out
	tidy.Env = blankEnv()
	if err := tidy.Run(); err != nil {
		return err
	}

	test := exec.Command("go", "build", "-o", "/tmp/runsc", "--mod=readonly", "./runsc")
	test.Stdout = os.Stdout
	test.Stderr = os.Stderr
	test.Dir = out
	test.Env = blankEnv()
	if err := test.Run(); err != nil {
		return err
	}

	if err := copyDataFiles(out); err != nil {
		return err
	}

	if err := switchRoot(out); err != nil {
		return err
	}

	return nil
}

func switchRoot(out string) error {
	out = filepath.Clean(out)
	fis, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}
	for _, fi := range fis {
		if fi.Name() != ".git" && fi.Name() != out {
			if err := os.RemoveAll(filepath.Join(".", fi.Name())); err != nil {
				return err
			}
		}
	}

	mv := exec.Command("sh", "-c", "mv "+out+"/* .")
	mv.Stdout = os.Stdout
	mv.Stderr = os.Stderr
	if err := mv.Run(); err != nil {
		return err
	}

	return os.RemoveAll(out)
}

func chmod(d string) error {
	return filepath.Walk(d, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return os.Chmod(p, 0755)
		}
		return os.Chmod(p, 0644)
	})
}

func copyDataFiles(out string) error {
	return filepath.Walk(".", func(p string, info os.FileInfo, err error) error {
		if p == ".git" || p == filepath.Clean(out) {
			return filepath.SkipDir
		}
		if info.IsDir() {
			return nil
		}
		base := filepath.Base(p)

		if base == "LICENSE" || strings.HasSuffix(base, ".md") || strings.HasSuffix(base, ".png") || strings.HasSuffix(base, ".svg") || base == ".gitignore" {
			if err := os.MkdirAll(filepath.Join(out, filepath.Dir(p)), 0701); err != nil {
				return err
			}
			dst, err := os.Create(filepath.Join(out, p))
			if err != nil {
				return err
			}
			src, err := os.Open(p)
			if err != nil {
				return err
			}
			if _, err := io.Copy(dst, src); err != nil {
				return err
			}
			dst.Close()
			src.Close()
			fmt.Printf("copied %s to %s\n", p, filepath.Join(out, p))
			return nil
		}
		return nil
	})
}

func fixNeedsAssembly(d string) error {
	return filepath.Walk(d, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		fis, err := ioutil.ReadDir(p)
		if err != nil {
			return err
		}
		hasAssembly := false
		needsAssembly := false
		for _, fi := range fis {
			name := fi.Name()
			if strings.HasSuffix(name, ".s") {
				hasAssembly = true
				break
			}
			if !strings.HasSuffix(name, ".go") {
				continue
			}
			dt, err := ioutil.ReadFile(filepath.Join(p, name))
			if err != nil {
				return err
			}
			if bytes.Contains(dt, []byte("go:linkname")) && bytes.Contains(dt, []byte("unsafe")) {
				needsAssembly = true
			}
		}
		if needsAssembly && !hasAssembly {
			dest := filepath.Join(p, "force_golinkname.s")
			fmt.Printf("write assembly %s\n", dest)
			return ioutil.WriteFile(dest, []byte{}, 0600)
		}
		return nil
	})
}

func readWorkspaceDeps() (deps []dep, err error) {
	dt, err := ioutil.ReadFile("WORKSPACE")
	if err != nil {
		return nil, err
	}
	r := regexp.MustCompile("importpath\\s+=\\s+\"([^\"]+)\",\\s+commit\\s+=\\s+\"([^\"]+)\"")
	for _, res := range r.FindAllStringSubmatch(string(dt), -1) {
		deps = append(deps, dep{Name: res[1], SHA: res[2]})
	}

	r = regexp.MustCompile("commit\\s+=\\s+\"([^\"]+)\",\\s+importpath\\s+=\\s+\"([^\"]+)\"")
	for _, res := range r.FindAllStringSubmatch(string(dt), -1) {
		deps = append(deps, dep{Name: res[2], SHA: res[1]})
	}

	return deps, nil
}

func readProtoDeps() (deps []dep, err error) {
	dt, err := exec.Command("bazel", "info", "output_base").Output()
	if err != nil {
		return nil, err
	}

	fp := filepath.Join(strings.TrimSpace(string(dt)), "external/io_bazel_rules_go/go/private/repositories.bzl")

	dt, err = ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile("remote\\s+=\\s+\"([^\"]+)\",\\s+commit\\s+=\\s+\"([^\"]+)\"")
	for _, res := range r.FindAllStringSubmatch(string(dt), -1) {
		switch res[1] {
		case "https://github.com/google/go-genproto":
			deps = append(deps, dep{Name: "google.golang.org/genproto", SHA: res[2]})
		case "https://github.com/golang/protobuf":
			deps = append(deps, dep{Name: "github.com/golang/protobuf", SHA: res[2]})
		}
		// fmt.Println(res[1], res[2])
	}
	return deps, nil
}

type dep struct {
	Name string
	SHA  string
}

func blankEnv() []string {
	return []string{"PATH=" + os.Getenv("PATH"), "HOME=" + os.Getenv("HOME")}
}
