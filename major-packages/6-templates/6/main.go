package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{Nome: "Curso de Go", CargaHoraria: 40},
		{Nome: "Curso de Python", CargaHoraria: 20},
		{Nome: "Curso de Java", CargaHoraria: 30},
	})
	if err != nil {
		panic(err)
	}
}
