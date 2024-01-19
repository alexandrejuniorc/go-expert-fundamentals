package main

import "fmt"

func main() {
	defer fmt.Println("Primeira Linha") // o defer deixa essa execução por último, muito utilizado em fechamento de conexões
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")

}
