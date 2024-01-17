package main

import "fmt"

func main() {
	var minhaVar interface{} = "Alexandre Junior"
	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %d e o valor de ok é %t\n", res, ok)

	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %d\n", res2) // como a resposta do ok foi falso, o compilador vai retornar um erro
}
