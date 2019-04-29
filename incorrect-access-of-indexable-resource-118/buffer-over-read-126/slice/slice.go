package main

import "log"

func main() {
	int_val := 4
	arr_val := [...]int{1, 2, 3}
	log.Println(arr_val[0:int_val])
}
