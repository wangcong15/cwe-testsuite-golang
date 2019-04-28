package main

import (
	"github.com/wangcong15/goassert"
)

func main() {
	var int_val int = 0
	goassert.AssertNEq(int_val, 0)
	int_val = 5 / int_val
}
