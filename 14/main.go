package main

import "fmt"

type Cliente struct {
	Nome string
}

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	fmt.Println(c.saldo)
	return c.saldo
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c Cliente) andou() {
	c.Nome = "Alexandre Junior"
	fmt.Printf("O cliente %s andou\n", c.Nome)
}

func main() {
	ale := Cliente{
		Nome: "Alexandre",
	}

	ale.andou()
	fmt.Printf("O valor da struct com nome é %v\n", ale.Nome)

	conta := Conta{
		saldo: 100,
	}

	conta.simular(200)
	fmt.Println(conta.saldo)
}
