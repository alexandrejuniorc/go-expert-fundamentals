package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("content.html").ParseFiles("content.html"))
	err := t.Execute(os.Stdout, Cursos{
		{Nome: "Curso de Go", CargaHoraria: 40},
		{Nome: "Curso de Python", CargaHoraria: 20},
		{Nome: "Curso de Java", CargaHoraria: 30},
	})
	if err != nil {
		panic(err)
	}
}
