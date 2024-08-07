package main

import (
	"fmt"
)

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

	// Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	println(r)

	ls := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(ls)

	// Slice defaults
	sd := []int{2, 3, 5, 7, 11, 13}

	sd = sd[0:4]
	fmt.Println(sd)

	sd = sd[:3] // O padrão é zero para o limite inferior
	fmt.Println(sd)

	sd = sd[1:] // comprimento do slice para o limite superior.
	fmt.Println(sd)

	// Slice length and capacity
	sl := []int{2, 3, 5, 7, 11, 13}
	printSlice(sl)

	// Slice the slice to give it zero length.
	sl = sl[:0]
	printSlice(sl)

	// Extend its length.
	sl = sl[:4]
	printSlice(sl)

	// Drop its first two values.
	sl = sl[2:]
	printSlice(sl)

	// slice nil
	var sn []int
	fmt.Println(sn, len(sn), cap(sn))
	if sn == nil {
		fmt.Println("nil!")
	}

	// slice with make
	am := make([]int, 5)
	printSliceMake("am", am)

	bm := make([]int, 0, 5)
	printSliceMake("bm", bm)

	cm := bm[:2]
	printSliceMake("cm", cm)

	dm := bm[2:5]
	printSliceMake("dm", dm)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
