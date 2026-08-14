package main

import (
	f "fmt"
	"strconv"
)

func main() {

	f.Println("Teste ", string(98))

	f.Println("Float ", float64(97))

	f.Println("Teste ", int(97))
	
	// int para string
	f.Println("Numero ", strconv.Itoa(93))
	
	// string para int
	num, erro := strconv.Atoi("93")
	f.Println("Numero ", num)
	f.Println("error ", erro)

}