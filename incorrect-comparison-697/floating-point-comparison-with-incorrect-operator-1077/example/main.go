package main

import (
	"log"
	"math"

	"github.com/wangcong15/goassert"
)

const prec_loss = 0.000001

func main() {
	x := 74.96
	y := 20.48
	b := x - y
	log.Println(b) // 54.47999999999999
	a := 54.48

	goassert.AssertGt(math.Abs(a-b), prec_loss)
	log.Println(a == b) // false
}
