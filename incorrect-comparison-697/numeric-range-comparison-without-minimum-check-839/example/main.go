package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/wangcong15/goassert"
)

func getSliceLength() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100) - 50
}

func main() {
	randLen := getSliceLength()
	if randLen > 50 {
		panic("Length Invalid")
	}
	var int_arr []int
	// randLen: numeric range comparison without minimum check
	goassert.AssertGt(randLen, 0)
	for i := 0; i < randLen; i++ {
		int_arr = append(int_arr, i)
	}
	log.Printf("Length of Array: %v", len(int_arr))
}
