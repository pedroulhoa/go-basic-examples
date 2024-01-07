package main

import "fmt"

var c, python, java bool

var x, y int = 1, 2

func main() {
	var i int

	var value1, value2, value3 = true, false, "String!"

	shortDeclaration1 := 3
	shortDeclaration2, shortDeclaration3 := true, "Example!"

	// 0, false, false, false
	fmt.Println(i, c, python, java)

	// 1, 2
	fmt.Println(x, y)

	// true, false, "String!"
	fmt.Println(value1, value2, value3)

	fmt.Println(shortDeclaration1, shortDeclaration2, shortDeclaration3)
}
