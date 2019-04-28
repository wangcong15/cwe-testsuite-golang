package main

import (
	"log"

	"github.com/wangcong15/goassert"
)

func main() {
	var int_val int = 128
	var int8_val int8
	int8_val = int8(int_val)
	log.Printf("%v -> %v", int_val, int8_val)
	goassert.AssertValEq(int_val, int8_val)
}
