package main

import (
	"log"
	"unsafe"
)

func main() {
	int_val := [...]int{1, 2, 3}
	addr := uintptr(unsafe.Pointer(&int_val))
	addr += 8
	int_p := (*int)(unsafe.Pointer(addr))
	log.Println(int_val)
	*int_p = 20
	log.Println(int_val)
}
