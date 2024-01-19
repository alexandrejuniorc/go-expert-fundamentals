package main

import "fmt"

func main() {
	// memória -> endereço -> valor
	a := 10         // memória
	fmt.Println(a)  // valor
	fmt.Println(&a) // endereço

	var ponteiro *int = &a
	*ponteiro = 20
	// fmt.Println(ponteiro)
	// fmt.Println(a)

	b := &a
	*b = 30
	fmt.Println(*b) // recebe o valor de b
}
