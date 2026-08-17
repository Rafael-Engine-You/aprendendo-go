package main

import "fmt"

func main() {
	var notas [3]float64

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	total := 0.0

	i := 0
	// Simulando um while
	for i < len(notas) {
		total += notas[i]
		i++ // i = i + 1
	}

	fmt.Printf("Total %.2f\n", total)
}