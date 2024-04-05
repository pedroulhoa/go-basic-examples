package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	//for
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// for continued / For is Go's "while"
	sumContinued := 1
	for sumContinued < 1000 {
		sumContinued += sumContinued
	}

	fmt.Println(sumContinued)

	// If
	fmt.Println(sqrt(2), sqrt(-4))
}
