package main

import (
	"fmt"
)

func main() {

	// array := [10]int {}
	// slices := []int {}
	maps := map[string]int { // CHAVE: VALOR
		"joao": 20,
		"antonio": 24,
		"maria": 30,
		"gustavo": 11,
	}

	fmt.Printf("Antonio tem %d anos\n", maps["antonio"])
	fmt.Printf("Tamanho do Map %d\n", len(maps))
}