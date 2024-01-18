package main

func main() {
	// no go existe apenas o FOR para percorrer uma sequência de valores
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}

	// utilizando um _ você não mostra a variável tanto em key como em value
	for key, value := range numeros {
		println(key, value)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// loop infinito, CUIDADO
	// utilizado pra consumir mensagem de uma fila, por exemplo, ou de uma coisa que nunca vai parar
	// for {
	// 	println("Hello, World!")
	// }
}
