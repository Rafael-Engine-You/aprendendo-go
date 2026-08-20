package main

import (
	"fmt"
	// c "aprendendo-go/calc"
	"time"
	cli "aprendendo-go/clientes"
	prod "aprendendo-go/produtos"
)

func main() {
	// fmt.Println("Somar7 10 + 5 = ",	 c.Soma(10,5))
	// fmt.Println("Subtrair 10 - 5 = ",c.Subtrair(10,7))

	dtNascimento, _ := time.Parse("02/06/2006", "24/05/1982")

	antonio := cli.Cliente {
		NomeCompleto : "Antônio Oliveira",
		DataNascimento: dtNascimento, 
		Cpf: "000.666.888-55",
		Contatos: []string {
			"86 9.8855.6633",
		 	"86 9.9988.5566",
		},
	}

	fmt.Println("Cliente: ", antonio.NomeCompleto)
	fmt.Printf("Data Nascimento: %s\n", 
				antonio.DataNascimento.Format("02/01/2006"))
	celular := prod.Produto {
		Nome: "Iphone 17 Pro Max",
		Preco: 16999.99,
		quantidade: 10

		fmt.Println("Pruduto : ", celular.Nome)
	}
}