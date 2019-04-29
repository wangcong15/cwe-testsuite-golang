package main

import (
	"log"

	"github.com/wangcong15/goassert"
)

func getUserAge(userId int) (age interface{}) {
	if userId > 0 {
		age = userId
	}
	return
}

func main() {
	age := getUserAge(-1)
	goassert.AssertNNil(age)
	log.Println(age)
}
