package main

import "fmt"

func main() {
	salarios := map[string]int{"Alexandre": 1000, "João": 2000, "Wesley": 3000}
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Wesley"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}
	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}

}
