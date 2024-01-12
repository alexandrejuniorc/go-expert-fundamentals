package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco // or
	// Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main() {
	ale := Cliente{
		Nome:  "Alexandre",
		Idade: 22,
		Ativo: true,
	}
	ale.Ativo = false
	ale.Address.Cidade = "Guaxupé" // or
	// ale.Endereco.Cidade = "São Paulo"

	ale.Desativar()
}
