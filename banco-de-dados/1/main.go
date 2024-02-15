package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

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

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Notebook", 1899.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	statement, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)") // ? é um placeholder para evitar SQL Injection
	if err != nil {
		return err
	}

	defer statement.Close() // é importante fechar o statement após a execução

	_, err = statement.Exec(product.ID, product.name, product.price) // executa a query com os valores passados
	if err != nil {
		return err
	}

	return nil
}
