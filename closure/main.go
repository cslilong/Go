// closure project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	var j int = 5
	a := func() func() {
		var i int = 10

		return func() {
			j++
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	a()

	j += 2

	a()
}
