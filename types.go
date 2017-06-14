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

// SELboolean as declared in selinux/selinux.h:305
type SELboolean struct {
	Name           []byte
	Value          int32
	ref4c7793c1    *C.SELboolean
	allocs4c7793c1 interface{}
}
