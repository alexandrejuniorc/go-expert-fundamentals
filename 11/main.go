package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {}

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

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
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
	Desativacao(ale)

	minhaEmpresa := Empresa{Nome: "Piramide"}
	Desativacao(minhaEmpresa) // or
	// minhaEmpresa.Desativar()

}
