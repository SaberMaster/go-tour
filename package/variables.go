package main

import "fmt"

var c, python, java bool

// init
var i, j int = 1, 2

// short var

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no"
	fmt.Println(i, j, c ,python, java)


	// only used in function body
	k := 3
	perl, ruby, rocket := false, true, "yes"
	fmt.Println(k, perl, ruby, rocket)
}

