package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	result, err := json.Marshal(conta) // convertendo para json guardando o resultado em result e guardando ele na variável
	if err != nil {
		panic(err)
	}
	println(string(result))

	err = json.NewEncoder(os.Stdout).Encode(conta) // convertendo para json e imprimindo no terminal e já entregando o json e não guardando ele
	if err != nil {
		panic(err)
	}

	//jsonPuro := []byte(`{"numero": 2, "saldo": 200}`) // posso usar das duas formas, se a struct tiver tags
	jsonPuro := []byte(`{"n": 2, "s": 200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX) // convertendo json para struct e guardando na variável
	if err != nil {
		panic(err)
	}
	print(contaX.Saldo)
}
