package main

import (
	"fmt"

	"github.com/wangcong15/goassert"
)

const PASSED = 1
const FAILED = 0
const UNKNOWN = 2

func security_check() int {
	return UNKNOWN // this value is beyond switch cases
}

func main() {

	status := security_check()
	valid_set := []int{PASSED, FAILED}
	goassert.AssertIntIn(status, valid_set)
	switch status {
	case PASSED:
		fmt.Println("PASSED")
	case FAILED:
		fmt.Println("FAILED")
	}
}
