package main

import "fmt"

type Atleta struct {
	nome string
	idade int
	peso float64
	altura float64
	sexo string
}

func main() {

	pedro := Atleta {
		nome : "Pedro Cardoso",
		idade : 25,
		peso : 60.0,
		altura : 1.72,
		sexo : "Masculino",
	}

	maria := Atleta{
		nome : "Maria Clara",
		idade : 20,
		peso : 45.0,
		altura : 1.62,
		sexo : "Feminino",	
	}

	atletas := []Atleta {pedro, maria}

	for i := 0; i < len(atletas); i++ {
		fmt.Println(atletas[i].nome)
	}

}