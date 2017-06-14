// MIT

// WARNING: This file has automatically been generated on Wed, 14 Jun 2017 13:41:12 +03.
// By https://git.io/c-for-go. DO NOT EDIT.

package libselinux

/*
#cgo pkg-config: libselinux
#include "selinux/selinux.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocSELbooleanMemory allocates memory for type C.SELboolean in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSELbooleanMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSELbooleanValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSELbooleanValue = unsafe.Sizeof([1]C.SELboolean{})

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *SELboolean) Ref() *C.SELboolean {
	if x == nil {
		return nil
	}
	return x.ref4c7793c1
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *SELboolean) Free() {
	if x != nil && x.allocs4c7793c1 != nil {
		x.allocs4c7793c1.(*cgoAllocMap).Free()
		x.ref4c7793c1 = nil
	}
}

// NewSELbooleanRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewSELbooleanRef(ref unsafe.Pointer) *SELboolean {
	if ref == nil {
		return nil
	}
	obj := new(SELboolean)
	obj.ref4c7793c1 = (*C.SELboolean)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *SELboolean) PassRef() (*C.SELboolean, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref4c7793c1 != nil {
		return x.ref4c7793c1, nil
	}
	mem4c7793c1 := allocSELbooleanMemory(1)
	ref4c7793c1 := (*C.SELboolean)(mem4c7793c1)
	allocs4c7793c1 := new(cgoAllocMap)
	var cname_allocs *cgoAllocMap
	ref4c7793c1.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Name)).Data)), cgoAllocsUnknown
	allocs4c7793c1.Borrow(cname_allocs)

	var cvalue_allocs *cgoAllocMap
	ref4c7793c1.value, cvalue_allocs = (C.int)(x.Value), cgoAllocsUnknown
	allocs4c7793c1.Borrow(cvalue_allocs)

	x.ref4c7793c1 = ref4c7793c1
	x.allocs4c7793c1 = allocs4c7793c1
	return ref4c7793c1, allocs4c7793c1

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x SELboolean) PassValue() (C.SELboolean, *cgoAllocMap) {
	if x.ref4c7793c1 != nil {
		return *x.ref4c7793c1, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *SELboolean) Deref() {
	if x.ref4c7793c1 == nil {
		return
	}
	hxfc4425b := (*sliceHeader)(unsafe.Pointer(&x.Name))
	hxfc4425b.Data = uintptr(unsafe.Pointer(x.ref4c7793c1.name))
	hxfc4425b.Cap = 0x7fffffff
	// hxfc4425b.Len = ?

	x.Value = (int32)(x.ref4c7793c1.value)
}

// safeString ensures that the string is NULL-terminated, a NULL-terminated copy is created otherwise.
func safeString(str string) string {
	if len(str) > 0 && str[len(str)-1] != '\x00' {
		str = str + "\x00"
	} else if len(str) == 0 {
		str = "\x00"
	}
	return str
}

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	str = safeString(str)
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// allocPCharMemory allocates memory for type *C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPCharValue = unsafe.Sizeof([1]*C.char{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackArgSSByte transforms a sliced Go data structure into plain C format.
func unpackArgSSByte(x [][]byte) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.char)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}

// packSSByte reads sliced Go data structure out from plain C format.
func packSSByte(v [][]byte, ptr0 **C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.char)(unsafe.Pointer(ptr0)))[i0]
		hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxf95e7c8.Data = uintptr(unsafe.Pointer(ptr1))
		hxf95e7c8.Cap = 0x7fffffff
		// hxf95e7c8.Len = ?
	}
}

// unpackArgSSELboolean transforms a sliced Go data structure into plain C format.
func unpackArgSSELboolean(x []SELboolean) (unpacked *C.SELboolean, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.SELboolean) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSELbooleanMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.SELboolean)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0], _ = x[i0].PassValue()
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.SELboolean)(unsafe.Pointer(h.Data))
	return
}

// packSSELboolean reads sliced Go data structure out from plain C format.
func packSSELboolean(v []SELboolean, ptr0 *C.SELboolean) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSELbooleanValue]C.SELboolean)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewSELbooleanRef(unsafe.Pointer(&ptr1))
	}
}

// allocPPCharMemory allocates memory for type **C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPPCharValue = unsafe.Sizeof([1]**C.char{})

// unpackArgSSSByte transforms a sliced Go data structure into plain C format.
func unpackArgSSSByte(x [][][]byte) (unpacked ***C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(****C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]**C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocPCharMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]*C.char)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			h := (*sliceHeader)(unsafe.Pointer(&x[i0][i1]))
			v1[i1] = (*C.char)(unsafe.Pointer(h.Data))
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (**C.char)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (***C.char)(unsafe.Pointer(h.Data))
	return
}

// packSSSByte reads sliced Go data structure out from plain C format.
func packSSSByte(v [][][]byte, ptr0 ***C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]**C.char)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfPtr]*C.char)(unsafe.Pointer(ptr1)))[i1]
			hxff2234b := (*sliceHeader)(unsafe.Pointer(&v[i0][i1]))
			hxff2234b.Data = uintptr(unsafe.Pointer(ptr2))
			hxff2234b.Cap = 0x7fffffff
			// hxff2234b.Len = ?
		}
	}
}
