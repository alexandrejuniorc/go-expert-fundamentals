package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	// tamanho, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()

	// leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// leitura de pouco em pouco, abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt") // abrindo o arquivo
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2) // leitura de pouco em pouco
	buffer := make([]byte, 3)           // buffer para leitura de pouco em pouco

	// loop para leitura de pouco em pouco
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	arquivo2.Close()

	// removendo um arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
