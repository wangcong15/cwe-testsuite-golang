package main

import (
	"log"

	"github.com/wangcong15/goassert"
)

func unTrustedNumber() int {
	return 0
}

func main() {
	x := unTrustedNumber()
	goassert.AssertGt(x, 0)
	y := make([]int, x)
	log.Println(y)
}
