package main

import (
	"fmt"
	"math"

	"github.com/wangcong15/goassert"
)

func main() {
	size := 1000000000000000
	goassert.AssertLt(size, math.MaxInt32)
	r := make([]int, size)

	fmt.Println("Size: ", len(r))
}
