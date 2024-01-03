package main

import "fmt"

// A function can take zero or more arguments.
// Notice that the type comes after the variable name.
func add(x int, y int) int {
	return x + y
}

// A function can return multiple results
func swap(textA string, textB string) (string, string) {
	return textA, textB
}

// shortened define type
// func add(x, y int)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(40, 40))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(split(10))
}
