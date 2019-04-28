package main

import (
	"math"

	"github.com/wangcong15/goassert"
)

func main() {
	var int8_val int8 = 127
	goassert.AssertLte(int16(int8_val)+1, math.MaxInt8)
	int8_val = int8_val + 1
}
