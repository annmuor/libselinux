// MIT

// WARNING: This file has automatically been generated on Thu, 08 Jun 2017 11:16:07 UTC.
// By https://git.io/c-for-go. DO NOT EDIT.

package libselinux

/*
#cgo pkg-config: libselinux
#include "selinux/selinux.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// security_context_t type as declared in selinux/selinux.h:17
type security_context_t []byte

// security_class_t type as declared in selinux/selinux.h:126
type security_class_t uint16

// security_class_mapping as declared in selinux/selinux.h:356
type security_class_mapping struct {
	name           string
	perms          [33]string
	ref2040d297    *C.struct_security_class_mapping
	allocs2040d297 interface{}
}
