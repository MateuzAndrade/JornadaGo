package main

import "fmt"

type pc int

var x pc

func main() {

	fmt.Printf("%v\n%T\n", x, x)
	x = 42
	fmt.Print(x)
}
