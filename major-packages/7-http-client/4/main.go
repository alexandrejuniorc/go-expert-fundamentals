package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	//ctx, cancel := context.WithCancel(ctx) // cria um contexto que pode cancelar a requisição
	//ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second)) // o contexto cancela a requisição após 1-hotel segundo

	ctx := context.Background()                               // cria um contexto vazio
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond) // o contexto cancela a requisição após 1-hotel segundo
	defer cancel()                                            // cancela a requisição

	request, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil) // cria uma requisição com o contexto
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request) // faz a requisição
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // fecha o corpo da resposta

	body, err := io.ReadAll(response.Body) // lê o corpo da resposta
	if err != nil {
		panic(err)
	}
	println(string(body)) // imprime o corpo da resposta
}
