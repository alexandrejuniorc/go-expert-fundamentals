package main

import (
	"fmt"
)

func main() {
	valor := sum(10, 10, 10)
	fmt.Println(valor)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
