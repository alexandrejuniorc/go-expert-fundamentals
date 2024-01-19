package main

import (
	"fmt"

	"github.com/alexandrejuniorc/matematica"
	"github.com/google/uuid"
)

// a lib uuid não foi importada automaticamente - executei go mod tidy para baixar as dependências e remover as dependencias não utilizadas

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %d\n", soma)

	fmt.Println(matematica.A) // pra utilizar variáveis ou funções de outro package ele deve começar com uma letra maiúscula

	carro := matematica.Carro{Marca: "Fiat"} // as propriedades de um struct devem iniciar com letra maiúscula para serem reconhecidas
	fmt.Println(carro.Marca)
	fmt.Println(carro.Andar()) // métodos das structs deve iniciar com letra minúscula

	fmt.Println(uuid.New())
}
