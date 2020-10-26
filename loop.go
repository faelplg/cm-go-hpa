package main

import (
	"fmt"
	"math"
)

func loop(x float64) float64 {
	for i:=0; i<1000000; i++ {
		x += math.Sqrt(x)
		fmt.Println("Code.Education Rocks!")
	}
	return x
}

func main() {
	x := 0.0001
	xRes := loop(x)
	fmt.Printf("%f", xRes)
}
