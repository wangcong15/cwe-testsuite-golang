package main

import (
	"github.com/wangcong15/goassert"
)

func unTrustedNumber() int16 {
	return -32768
}

func main() {
	var int16_val int16 = unTrustedNumber()
	goassert.AssertUnderflow(int16_val, 1, int16_val-1)
	int16_val = int16_val - 1
}
