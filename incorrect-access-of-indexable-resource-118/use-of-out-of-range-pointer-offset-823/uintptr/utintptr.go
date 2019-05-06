package main

import (
	"unsafe"

	"github.com/wangcong15/goassert"
)

func main() {
	int_val := [...]int{1, 2, 3}
	addr := uintptr(unsafe.Pointer(&int_val))
	addr += 30
	int_p := (*int)(unsafe.Pointer(addr))
	goassert.AssertLte(uintptr(unsafe.Pointer(int_p))-uintptr(unsafe.Pointer(&int_val)), len(int_val)*int(unsafe.Sizeof(int_val[0])))
	*int_p = 20
}
