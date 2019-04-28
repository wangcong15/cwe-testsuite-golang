package main

import (
	"log"

	"github.com/wangcong15/goassert"
)

func main() {
	var int_val int = 2147483648
	var int32_val int32
	int32_val = int32(int_val)
	log.Printf("%v -> %v", int_val, int32_val)
	goassert.AssertValEq(int_val, int32_val)
}
