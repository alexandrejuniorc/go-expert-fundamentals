package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	ale := Cliente{
		Nome:  "Alexandre",
		Idade: 22,
		Ativo: true,
	}
	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", ale.Nome, ale.Idade, ale.Ativo)

	ale.Ativo = false
	fmt.Println(ale.Ativo)
}
