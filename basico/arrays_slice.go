package main

import (
	"fmt"
	"slices"
)

func main() {
	
	// numeros := [5]int {10, 20, 30, 40} // Arrays 5-1
	// numeros[4] =50
	// fmt.Println("Tamanho", numeros)
	// fmt.Println("Tamanho", len(numeros))
	
	frutas := []string {"Goaiba", "Maçã", "Abacaxi" } //slice
	frutas = append(frutas, "Pera")
	//slices.Sort(frutas)
	// slices.Reverse(frutas)
	frutas = slices.Delete(frutas, 1, 3)
	fmt.Println(frutas)
	fmt.Println("Tamanho da lista: ", len(frutas))
	fmt.Println("Tem maçã na lista? ", slices.Contains(frutas, "Maçã"))

}