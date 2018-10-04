package time

import (
	"fmt"
	ssync "gvisor.googlesource.com/gvisor/pkg/sync"
	"reflect"
	"strings"
	"unsafe"
)

// SeqAtomicLoad returns a copy of *ptr, ensuring that the read does not race
// with any writer critical sections in sc.
func SeqAtomicLoadParameters(sc *ssync.SeqCount, ptr *Parameters) Parameters {
	// This function doesn't use SeqAtomicTryLoad because doing so is
	// measurably, significantly (~20%) slower; Go is awful at inlining.
	var val Parameters
	for {
		epoch := sc.BeginRead()
		if ssync.RaceEnabled {

			ssync.Memmove(unsafe.Pointer(&val), unsafe.Pointer(ptr), unsafe.Sizeof(val))
		} else {

			val = *ptr
		}
		if sc.ReadOk(epoch) {
			break
		}
	}
	return val
}

// SeqAtomicTryLoad returns a copy of *ptr while in a reader critical section
// in sc initiated by a call to sc.BeginRead() that returned epoch. If the read
// would race with a writer critical section, SeqAtomicTryLoad returns
// (unspecified, false).
func SeqAtomicTryLoadParameters(sc *ssync.SeqCount, epoch ssync.SeqCountEpoch, ptr *Parameters) (Parameters, bool) {
	var val Parameters
	if ssync.RaceEnabled {
		ssync.Memmove(unsafe.Pointer(&val), unsafe.Pointer(ptr), unsafe.Sizeof(val))
	} else {
		val = *ptr
	}
	return val, sc.ReadOk(epoch)
}

func initParameters() {
	var val Parameters
	typ := reflect.TypeOf(val)
	name := typ.Name()
	if ptrs := ssync.PointersInType(typ, name); len(ptrs) != 0 {
		panic(fmt.Sprintf("SeqAtomicLoad<%s> is invalid since values %s of type %s contain pointers:\n%s", typ, name, typ, strings.Join(ptrs, "\n")))
	}
}
