package main

import (
	"log"
	"reflect"

	"github.com/wangcong15/goassert"
)

type Haha interface {
}

type wow struct {
}

func main() {
	var s Haha = wow{}
	n := reflect.TypeOf(s)

	goassert.AssertEq(n.PkgPath(), "main")
	// sometimes we check the name, but ignore the consistency of package
	if n.Name() == "wow" {
		log.Println("This is the structure we want!")
	}
}
