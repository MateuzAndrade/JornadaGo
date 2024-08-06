package main

import "fmt"

type pc int

var x pc
var y int

func main() {

	fmt.Printf("%v\n%T\n", x, x)
	x = 42
	fmt.Print(x)

	fmt.Println("\n---------------X acima e Y abaixo------------------------")
	y = int(x)
	fmt.Printf("%v\n%T\n", y, y)

}
