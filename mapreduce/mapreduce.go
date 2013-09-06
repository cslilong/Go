package main

import (
	"fmt"
)

type hole chan int

func des(w hole, h int) {
	w <- h
}

func main() {
	n := 10
	w := make(hole, n)

	//map
	for i := 0; i < n; i++ {
		go des(w, i)
	}
	//reduce
	t := 0
	for i := 0; i < n; i++ {
		t += <-w
	}
	fmt.Println("total:", t)
}
