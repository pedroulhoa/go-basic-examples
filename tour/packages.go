// main package where the program starts
package main

import (
	"fmt"
	"math/rand"
)

// Is it possible to import separately
// import "fmt"
// import "math"

func main() {
	fmt.Println("My favorite number is: ", rand.Intn(10))
}
