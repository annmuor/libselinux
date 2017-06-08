// MIT

// WARNING: This file has automatically been generated on Thu, 08 Jun 2017 11:26:16 UTC.
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

// Getcon function as declared in selinux/selinux.h:29
func Getcon(Con [][]byte) int32 {
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.getcon(cCon)
	packSSByte(Con, cCon)
	__v := (int32)(__ret)
	return __v
}

// Setcon function as declared in selinux/selinux.h:40
func Setcon(Con string) int32 {
	Con = safeString(Con)
	cCon, _ := unpackPCharString(Con)
	__ret := C.setcon(cCon)
	runtime.KeepAlive(Con)
	__v := (int32)(__ret)
	return __v
}

// Getpidcon function as declared in selinux/selinux.h:45
func Getpidcon(Pid int32, Con [][]byte) int32 {
	cPid, _ := (C.pid_t)(Pid), cgoAllocsUnknown
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.getpidcon(cPid, cCon)
	packSSByte(Con, cCon)
	__v := (int32)(__ret)
	return __v
}

// Getprevcon function as declared in selinux/selinux.h:50
func Getprevcon(Con [][]byte) int32 {
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.getprevcon(cCon)
	packSSByte(Con, cCon)
	__v := (int32)(__ret)
	return __v
}

// Getexeccon function as declared in selinux/selinux.h:56
func Getexeccon(Con [][]byte) int32 {
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.getexeccon(cCon)
	packSSByte(Con, cCon)
	__v := (int32)(__ret)
	return __v
}

// Setexeccon function as declared in selinux/selinux.h:61
func Setexeccon(Con string) int32 {
	Con = safeString(Con)
	cCon, _ := unpackPCharString(Con)
	__ret := C.setexeccon(cCon)
	runtime.KeepAlive(Con)
	__v := (int32)(__ret)
	return __v
}

// Getfscreatecon function as declared in selinux/selinux.h:67
func Getfscreatecon(Con [][]byte) int32 {
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.getfscreatecon(cCon)
	packSSByte(Con, cCon)
	__v := (int32)(__ret)
	return __v
}

// Setfscreatecon function as declared in selinux/selinux.h:72
func Setfscreatecon(Context string) int32 {
	Context = safeString(Context)
	cContext, _ := unpackPCharString(Context)
	__ret := C.setfscreatecon(cContext)
	runtime.KeepAlive(Context)
	__v := (int32)(__ret)
	return __v
}

// Getkeycreatecon function as declared in selinux/selinux.h:78
func Getkeycreatecon(Con [][]byte) int32 {
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.getkeycreatecon(cCon)
	packSSByte(Con, cCon)
	__v := (int32)(__ret)
	return __v
}

// Setkeycreatecon function as declared in selinux/selinux.h:83
func Setkeycreatecon(Context string) int32 {
	Context = safeString(Context)
	cContext, _ := unpackPCharString(Context)
	__ret := C.setkeycreatecon(cContext)
	runtime.KeepAlive(Context)
	__v := (int32)(__ret)
	return __v
}

// Getsockcreatecon function as declared in selinux/selinux.h:89
func Getsockcreatecon(Con [][]byte) int32 {
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.getsockcreatecon(cCon)
	packSSByte(Con, cCon)
	__v := (int32)(__ret)
	return __v
}

// Setsockcreatecon function as declared in selinux/selinux.h:94
func Setsockcreatecon(Context string) int32 {
	Context = safeString(Context)
	cContext, _ := unpackPCharString(Context)
	__ret := C.setsockcreatecon(cContext)
	runtime.KeepAlive(Context)
	__v := (int32)(__ret)
	return __v
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

// Lgetfilecon function as declared in selinux/selinux.h:103
func Lgetfilecon(Path string, Con [][]byte) int32 {
	Path = safeString(Path)
	cPath, _ := unpackPCharString(Path)
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.lgetfilecon(cPath, cCon)
	packSSByte(Con, cCon)
	runtime.KeepAlive(Path)
	__v := (int32)(__ret)
	return __v
}

// Fgetfilecon function as declared in selinux/selinux.h:105
func Fgetfilecon(Fd int32, Con [][]byte) int32 {
	cFd, _ := (C.int)(Fd), cgoAllocsUnknown
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.fgetfilecon(cFd, cCon)
	packSSByte(Con, cCon)
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

// Lsetfilecon function as declared in selinux/selinux.h:111
func Lsetfilecon(Path string, Con string) int32 {
	Path = safeString(Path)
	cPath, _ := unpackPCharString(Path)
	Con = safeString(Con)
	cCon, _ := unpackPCharString(Con)
	__ret := C.lsetfilecon(cPath, cCon)
	runtime.KeepAlive(Con)
	runtime.KeepAlive(Path)
	__v := (int32)(__ret)
	return __v
}

// Fsetfilecon function as declared in selinux/selinux.h:113
func Fsetfilecon(Fd int32, Con string) int32 {
	cFd, _ := (C.int)(Fd), cgoAllocsUnknown
	Con = safeString(Con)
	cCon, _ := unpackPCharString(Con)
	__ret := C.fsetfilecon(cFd, cCon)
	runtime.KeepAlive(Con)
	__v := (int32)(__ret)
	return __v
}

// Getpeercon function as declared in selinux/selinux.h:120
func Getpeercon(Fd int32, Con [][]byte) int32 {
	cFd, _ := (C.int)(Fd), cgoAllocsUnknown
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.getpeercon(cFd, cCon)
	packSSByte(Con, cCon)
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

// Matchmediacon function as declared in selinux/selinux.h:489
func Matchmediacon(Media string, Con [][]byte) int32 {
	Media = safeString(Media)
	cMedia, _ := unpackPCharString(Media)
	cCon, _ := unpackArgSSByte(Con)
	__ret := C.matchmediacon(cMedia, cCon)
	packSSByte(Con, cCon)
	runtime.KeepAlive(Media)
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

// Setexecfilecon function as declared in selinux/selinux.h:603
func Setexecfilecon(Filename string, Fallback_type string) int32 {
	Filename = safeString(Filename)
	cFilename, _ := unpackPCharString(Filename)
	Fallback_type = safeString(Fallback_type)
	cFallback_type, _ := unpackPCharString(Fallback_type)
	__ret := C.setexecfilecon(cFilename, cFallback_type)
	runtime.KeepAlive(Fallback_type)
	runtime.KeepAlive(Filename)
	__v := (int32)(__ret)
	return __v
}
