package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		// Imprime no comand line do servidor
		log.Println("Request procesada com sucesso")

		// Imprime no navegador
		w.Write([]byte("Request procesada com sucesso"))

	case <-ctx.Done():
		// Imprime no comand line do servidor
		log.Println("Request cancelada pelo cliente")
	}
}
