package main

import "fmt"

func main() {
	//for
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// for continued
	sumContinued := 1
	for sumContinued < 1000 {
		sumContinued += sumContinued
	}

	fmt.Println(sumContinued)
}
