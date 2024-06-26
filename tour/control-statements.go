package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
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

	// If with a short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// switch no condition
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer
	defer fmt.Println("world")
	fmt.Println("Hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
