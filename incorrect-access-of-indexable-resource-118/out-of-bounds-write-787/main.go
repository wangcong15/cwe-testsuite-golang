package main

import (
	"log"
)

func unTrustFunc() int {
	return 4
}

func main() {
	arr := [...]int{1, 2, 3, 4}
	idx := unTrustFunc()
	arr[idx] = 5
	log.Println(arr)
}
