package main

import (
	"log"
)

func unTrustFunc() int {
	return 4
}

func main() {
	int_val := unTrustFunc()
	arr_val := [...]int{1, 2, 3}
	log.Println(arr_val[0:int_val])
}
