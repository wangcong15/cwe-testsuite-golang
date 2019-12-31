package main

import (
	"github.com/wangcong15/goassert"
	"log"
)

func main() {
	happyNewYear := []string{"Happy", "New", "Year"}
	lens := len(happyNewYear)
	goassert.AssertLt(lens, len(happyNewYear))
	for i := 0; i <= lens; i++ {
		log.Println(happyNewYear[i])
	}
}
