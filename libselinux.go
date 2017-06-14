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
	"unsafe"
)

// Freecon function as declared in selinux/selinux.h:20
func Freecon(Con []byte) {
	cCon, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&Con)).Data)), cgoAllocsUnknown
	C.freecon(cCon)
}

// Getfilecon function as declared in selinux/selinux.h:101
func Getfilecon(Path string, Con [][]byte) int32 {
	Path = safeString(Path)
	cPath, _ := unpackPCharString(Path)
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.getfilecon(cPath, cCon)
	packSSByte(Con, cCon)
	runtime.KeepAlive(Path)
	__v := (int32)(__ret)
	return __v
}

// Setfilecon function as declared in selinux/selinux.h:109
func Setfilecon(Path string, Con string) int32 {
	Path = safeString(Path)
	cPath, _ := unpackPCharString(Path)
	Con = safeString(Con)
	cCon, _ := unpackPCharString(Con)
	__ret := C.setfilecon(cPath, cCon)
	runtime.KeepAlive(Con)
	runtime.KeepAlive(Path)
	__v := (int32)(__ret)
	return __v
}

// Security_set_boolean_list function as declared in selinux/selinux.h:307
func Security_set_boolean_list(Boolcnt uint, Boollist []SELboolean, Permanent int32) int32 {
	cBoolcnt, _ := (C.size_t)(Boolcnt), cgoAllocsUnknown
	cBoollist, _ := unpackArgSSELboolean(Boollist)
	cPermanent, _ := (C.int)(Permanent), cgoAllocsUnknown
	__ret := C.security_set_boolean_list(cBoolcnt, cBoollist, cPermanent)
	packSSELboolean(Boollist, cBoollist)
	__v := (int32)(__ret)
	return __v
}

// Security_getenforce function as declared in selinux/selinux.h:326
func Security_getenforce() int32 {
	__ret := C.security_getenforce()
	__v := (int32)(__ret)
	return __v
}

// Security_setenforce function as declared in selinux/selinux.h:329
func Security_setenforce(Value int32) int32 {
	cValue, _ := (C.int)(Value), cgoAllocsUnknown
	__ret := C.security_setenforce(cValue)
	__v := (int32)(__ret)
	return __v
}

// Security_policyvers function as declared in selinux/selinux.h:338
func Security_policyvers() int32 {
	__ret := C.security_policyvers()
	__v := (int32)(__ret)
	return __v
}

// Security_get_boolean_names function as declared in selinux/selinux.h:341
func Security_get_boolean_names(Names [][][]byte, Len []int32) int32 {
	cNames, _ := unpackArgSSSByte(Names)
	cLen, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&Len)).Data)), cgoAllocsUnknown
	__ret := C.security_get_boolean_names(cNames, cLen)
	packSSSByte(Names, cNames)
	__v := (int32)(__ret)
	return __v
}

// Security_get_boolean_pending function as declared in selinux/selinux.h:344
func Security_get_boolean_pending(Name string) int32 {
	Name = safeString(Name)
	cName, _ := unpackPCharString(Name)
	__ret := C.security_get_boolean_pending(cName)
	runtime.KeepAlive(Name)
	__v := (int32)(__ret)
	return __v
}

// Security_get_boolean_active function as declared in selinux/selinux.h:347
func Security_get_boolean_active(Name string) int32 {
	Name = safeString(Name)
	cName, _ := unpackPCharString(Name)
	__ret := C.security_get_boolean_active(cName)
	runtime.KeepAlive(Name)
	__v := (int32)(__ret)
	return __v
}

// Matchpathcon_fini function as declared in selinux/selinux.h:447
func Matchpathcon_fini() {
	C.matchpathcon_fini()
}

// Matchpathcon function as declared in selinux/selinux.h:460
func Matchpathcon(Path string, Mode uint32, Con [][]byte) int32 {
	Path = safeString(Path)
	cPath, _ := unpackPCharString(Path)
	cMode, _ := (C.mode_t)(Mode), cgoAllocsUnknown
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.matchpathcon(cPath, cMode, cCon)
	packSSByte(Con, cCon)
	runtime.KeepAlive(Path)
	__v := (int32)(__ret)
	return __v
}

// Selinux_getpolicytype function as declared in selinux/selinux.h:511
func Selinux_getpolicytype(Policytype [][]byte) int32 {
	cPolicytype, _ := unpackArgSSByte(Policytype)
	__ret := C.selinux_getpolicytype(cPolicytype)
	packSSByte(Policytype, cPolicytype)
	__v := (int32)(__ret)
	return __v
}
