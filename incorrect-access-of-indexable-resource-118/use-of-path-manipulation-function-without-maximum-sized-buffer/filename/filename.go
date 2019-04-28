package main

import (
	"io/ioutil"
	"log"

	"github.com/wangcong15/goassert"
)

func main() {
	filename := []byte("This is a place holder")
	files, _ := ioutil.ReadDir("../..")
	for _, f := range files {
		goassert.AssertGte(len(filename), len(f.Name()))
		copy(filename, f.Name())
		s := string(filename)
		log.Println(s)
	}
}
