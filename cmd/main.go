package main

import (
	"fmt"

	doc "golang_study/cmd/documento"
)

func main() {
	var i = doc.Rg{
		Identificacao: doc.Identificacao{
			Nome:    "Fabiano",
			Emissao: "01-01-1992",
		},
	}

	CreateDocument(&i)
	var rg = i.Identificacao.Numero

	var c = doc.Cpf{
		Identificacao: doc.Identificacao{
			Nome:    "Janaina",
			Emissao: "05-04-1991",
		},
	}

	CreateDocument(&c)
	var cpf = c.Identificacao.Numero

	fmt.Println("rg:", rg, "Cpf:", cpf)
}

func CreateDocument(d doc.Operacoes) {
	d.Generate()
}
