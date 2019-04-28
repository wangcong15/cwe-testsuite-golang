package main

import (
	"math"

	"github.com/wangcong15/goassert"
)

func main() {
	var int16_val int16 = -32768
	goassert.AssertGte(int32(int16_val)-1, math.MinInt16)
	int16_val = int16_val - 1
}
