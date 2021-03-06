package main

import (
	"log"
)

func protect(g func()) {
	defer func() {
		log.Println("done")
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}

func main() {
	var s []byte
	protect(func() { s[0] = 0 })
	protect(func() { panic(42) })
	log.Println("main")
	s[0] = 42
}
