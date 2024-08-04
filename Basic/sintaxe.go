package main

import "fmt"

func main() {
	x := 10 //declaracao
	y := 15
	z := "Ola"
	x = 30 // atribuicao

	fmt.Println(x + y)
	fmt.Printf("z: %v,%T", z, z)
}
