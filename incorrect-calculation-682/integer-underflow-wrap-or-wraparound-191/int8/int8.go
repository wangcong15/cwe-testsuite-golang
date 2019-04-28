package main

import (
	"math"

	"github.com/wangcong15/goassert"
)

func main() {
	var int8_val int8 = -128
	goassert.AssertGte(int16(int8_val)-1, math.MinInt8)
	int8_val = int8_val - 1
}
