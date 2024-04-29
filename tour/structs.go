package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Structs literals
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v, v.X)

	// pointers in structs
	pointer := &v
	pointer.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
