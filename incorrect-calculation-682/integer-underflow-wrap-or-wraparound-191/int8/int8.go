package main

import "github.com/wangcong15/goassert"

func unTrustedNumber() int8 {
	return -128
}

func main() {
	var int8_val int8 = unTrustedNumber()
	goassert.AssertUnderflow(int8_val, 1, int8_val-1)
	int8_val = int8_val - 1
}
