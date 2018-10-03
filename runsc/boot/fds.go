// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package boot

import (
	"fmt"

	"gvisor.googlesource.com/gvisor/pkg/sentry/context"
	"gvisor.googlesource.com/gvisor/pkg/sentry/fs"
	"gvisor.googlesource.com/gvisor/pkg/sentry/fs/host"
	"gvisor.googlesource.com/gvisor/pkg/sentry/kernel"
	"gvisor.googlesource.com/gvisor/pkg/sentry/kernel/kdefs"
	"gvisor.googlesource.com/gvisor/pkg/sentry/limits"
)

// createFDMap creates an FD map that contains stdin, stdout, and stderr. If
// console is true, then ioctl calls will be passed through to the host FD.
// Upon success, createFDMap dups then closes stdioFDs.
func createFDMap(ctx context.Context, k *kernel.Kernel, l *limits.LimitSet, console bool, stdioFDs []int) (*kernel.FDMap, error) {
	if len(stdioFDs) != 3 {
		return nil, fmt.Errorf("stdioFDs should contain exactly 3 FDs (stdin, stdout, and stderr), but %d FDs received", len(stdioFDs))
	}

	fdm := k.NewFDMap()
	defer fdm.DecRef()

	// Maps sandbox FD to host FD.
	fdMap := map[int]int{
		0: stdioFDs[0],
		1: stdioFDs[1],
		2: stdioFDs[2],
	}
	mounter := fs.FileOwnerFromContext(ctx)

	for sfd, hfd := range fdMap {
		file, err := host.ImportFile(ctx, hfd, mounter, console /* isTTY */)
		if err != nil {
			return nil, fmt.Errorf("failed to import fd %d: %v", hfd, err)
		}
		defer file.DecRef()
		if err := fdm.NewFDAt(kdefs.FD(sfd), file, kernel.FDFlags{}, l); err != nil {
			return nil, fmt.Errorf("failed to add imported fd %d to FDMap: %v", hfd, err)
		}
	}

	fdm.IncRef()
	return fdm, nil
}
