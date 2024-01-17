package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Alexandre": 1000, "João": 2000, "Wesley": 3000}
	println(Soma(m))

	m2 := map[string]float64{"Alexandre": 100.20, "João": 2000.3, "Wesley": 300.0}
	println(Soma(m2))

	m3 := map[string]MyNumber{"Alexandre": 1000, "João": 2000, "Wesley": 3000}
	println(Soma(m3))

	println(Compara(10, 10.0))
	println(Compara(10, 20))
}
