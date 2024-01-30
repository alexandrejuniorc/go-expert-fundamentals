package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name":"John"}`))
	response, err := client.Post("https://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	io.CopyBuffer(os.Stdout, response.Body, nil)

}
