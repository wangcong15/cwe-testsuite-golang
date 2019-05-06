package main

import (
	"unsafe"

	"github.com/wangcong15/goassert"
)

func main() {
	int_val := 1
	addr := uintptr(unsafe.Pointer(&int_val))
	addr += 8
	int_p := (*int)(unsafe.Pointer(addr))
	goassert.AssertEq(int_p, &int_val)
	*int_p = 2
}
