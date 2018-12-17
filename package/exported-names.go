package main

import (
	"fmt"
)

// not exported
const pi = 1234

// exported
const PI = 1234


func main() {
	fmt.Println(pi)
}
