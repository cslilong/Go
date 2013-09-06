package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := new(big.Int)
	x.SetString("1", 10)
	y := new(big.Int)
	y.SetString("-2", 10)
	x.Add(x, y)
	fmt.Println(x)
	x.Mul(x, y)
	fmt.Println(x)
}
