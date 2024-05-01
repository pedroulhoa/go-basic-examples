package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:4]
	fmt.Println(s)

	var slice []int
	var temp = [2]int{1, 2}
	slice = temp[0:2]

	fmt.Println(slice)

	names := [4]string{
		"Peter",
		"Joe",
		"Paul",
		"John",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "TEST"
	fmt.Println(a, b)
	fmt.Println(names)
}
