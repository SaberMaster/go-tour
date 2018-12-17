package main

import "fmt"

const Pi = 3.14

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const World = "世界"
	const Truth = true
	fmt.Println("Hello", World, Pi, Truth)


	//fmt.Println(Big)
	//fmt.Println(Small)

	fmt.Println(needInt(Small))
	//fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
