package main

import (
	"log"
	"unsafe"
)

func main() {
	int_val := 1
	addr := uintptr(unsafe.Pointer(&int_val))
	addr += 8
	int_p := (*int)(unsafe.Pointer(addr))
	log.Println(int_p)
	*int_p = 2
	log.Println(*int_p)
}
