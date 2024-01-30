package main

import (
	"fmt"
	"math/cmplx"
)

var c, python, java bool

var x, y int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

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

	// test others types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Zero values
	var intzero int
	var floatzero float64
	var boolzero bool
	var stringzero string
	fmt.Printf("zero values: %v %v %v %q\n", intzero, floatzero, boolzero, stringzero)

	// type conversions
	var typeconversion int = 42
	var newfloat float64 = float64(typeconversion)
	var newuint uint = uint(newfloat)

	// simplified
	// i := 42
	// f := float64(i)
	// u := uint(f)

	fmt.Printf("type conversions: %v %v\n", newfloat, newuint)
}
