package main

import (
	"math/cmplx"
	"fmt"
	"math"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("%T, %v\n", ToBe, ToBe)
	fmt.Printf("%T, %v\n", MaxInt, MaxInt)
	fmt.Printf("%T, %v\n", z, z)

	// zero

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// type conversion
	var x, y int = 3, 4
	fl := math.Sqrt(float64(x*x + y*y))
	z := uint(fl)
	fmt.Println(z, y, z)

	// type inference
	v := 42

	fmt.Printf("%T", v)
}
