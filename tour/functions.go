package main

import "fmt"

// A function can take zero or more arguments.
// Notice that the type comes after the variable name.
func add(x int, y int) int {
	return x + y
}

// shortened
// func add(x, y int)

func main() {
	fmt.Println(add(40, 40))
}
