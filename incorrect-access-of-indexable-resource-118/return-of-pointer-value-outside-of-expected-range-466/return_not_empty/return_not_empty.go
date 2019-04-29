package main

import (
	"log"

	"github.com/wangcong15/goassert"
)

func getList(userId int) (list []int) {
	if userId > 0 {
		list = append(list, userId)
	}
	return
}

func main() {
	age := getList(-1)
	// convert into []interface{}
	s := make([]interface{}, len(age))
	for i, v := range age {
		s[i] = v
	}
	goassert.AssertNEmpty(s)
	log.Println(age)
}
