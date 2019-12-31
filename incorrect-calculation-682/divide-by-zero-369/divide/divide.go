package main

import (
	"log"
)

func unTrustFunc() int {
	return 0
}

func main() {
	var intVal int = unTrustFunc()
	intVal = 5 / intVal
	log.Println(intVal)
}
