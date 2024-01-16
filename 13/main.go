package main

import "fmt"

func soma(a, b *int) int {
	fmt.Println(*a, b)

	*a = 50
	*b = 50

	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 10
	soma(&minhaVar1, &minhaVar2)

	fmt.Println(minhaVar1)
	fmt.Println(minhaVar2)

}
