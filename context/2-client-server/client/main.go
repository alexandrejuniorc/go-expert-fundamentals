package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // Cria um contexto com timeout de 5 segundos
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil) // Cria a requisiçãos
	if err != nil {
		panic(err)
	}

	result, err := http.DefaultClient.Do(request) // Faz a requisição
	if err != nil {
		panic(err)
	}
	defer result.Body.Close()
	io.Copy(os.Stdout, result.Body) // Imprime no terminal
}
