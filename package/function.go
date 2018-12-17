package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// short
func add2(x, y int) int {
	return add(x, y)
}

// multi result

func swap(x, y string) (string, string) {
	return y, x
}

// named result
func split(sum int) (x, y int) {
	x = sum * 2 / 3
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))

	fmt.Println(add2(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	x, y := split(10)
	fmt.Println(x, y)
}
