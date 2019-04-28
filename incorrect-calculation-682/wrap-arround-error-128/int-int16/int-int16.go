package main

import (
	"log"

	"github.com/wangcong15/goassert"
)

func main() {
	var int_val int = 32768
	var int16_val int16
	int16_val = int16(int_val)
	log.Printf("%v -> %v", int_val, int16_val)
	goassert.AssertValEq(int_val, int16_val)
}
