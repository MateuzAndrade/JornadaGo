package main

import "fmt"

var x int = 5
var valida bool = false

func main() {

	fmt.Println("-----------------")
	fmt.Println("Valida IfElse \n")
	if x >= 18 {
		valida = true
		fmt.Println(valida)
	} else {
		fmt.Println(valida)
	}
	fmt.Println("-----------------")

}
