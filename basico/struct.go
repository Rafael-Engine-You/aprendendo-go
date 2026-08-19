package main

import "fmt"

type Conta struct {
	nome string
	saldo float64
	conta int
	agencia int
}

func main() {

	mariaConta := Conta {
		nome: "Maria Clara",
		saldo: 1000,
		conta: 100,
		agencia: 1,
	}	
	joseConta := Conta {
		nome : "José Antônio",
		saldo : 1020,
		conta : 100,
		agencia : 2,
	}

	contas := []Conta{ mariaConta, joseConta }

	for i := 0; i < len(contas); i++ {
		fmt.Printf("Cliente %s saldo de R$ %.2f\n", 
		contas[i].nome, contas[i].saldo)
	}

}