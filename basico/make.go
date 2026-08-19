package main

import "fmt"

func main() {

	idades :=  make([]int, 3)
	// idades := []int {}

	idades[0] = 10
	idades[1] = 20
	idades[2] = 30

	fmt.Println(idades)
	
	moedas := make(map[string]float64)

	moedas["dolar"] = 5.08
	moedas["real"]  = 1.0
	moedas["euro"]  = 7.50

	fmt.Println(moedas)
	


}