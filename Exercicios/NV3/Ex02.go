package main

import (
	"fmt"
	"time"
)

func main() {

	anoNascimento := 1995
	anoAtual := time.Now().Year()

	for anoNascimento <= anoAtual {
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}
