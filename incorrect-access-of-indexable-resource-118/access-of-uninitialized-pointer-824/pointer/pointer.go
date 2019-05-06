package main

import (
	"log"

	"github.com/wangcong15/goassert"
)

func main() {
	var a *int
	goassert.AssertNValEq(a, nil)
	log.Println(a)
}
