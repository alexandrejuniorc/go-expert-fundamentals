package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "ale"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 10

	fmt.Println(len(meuArray) - 1)         // last array value
	fmt.Println(len(meuArray))             // array length
	fmt.Println(meuArray[len(meuArray)-1]) // get last item on array
	fmt.Println(meuArray[0])               // get first item on array

	for index, value := range meuArray {
		fmt.Printf("O valor do indice é %v e o valor é %d\n", index, value)
	}
}
