package main

import "github.com/google/uuid"

type Product struct {
	ID    string
	name  string
	price float64
}

// NewProduct é uma função que cria um novo produto
func NewProduct(name string, price float64) *Product {
	// é importante que o ID seja gerado no momento da criação do produto
	return &Product{
		ID:    uuid.New().String(),
		name:  name,
		price: price,
	}
}

func main() {}
