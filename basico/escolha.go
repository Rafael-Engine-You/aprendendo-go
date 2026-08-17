package main

import (
	f "fmt"
	t "time"
)

func main() {
	var nome string
	t := t.Now()

	fmt.Print("Informe seu nome: ")
	fmt.Scan(&nome)

	switch {
		case t.Hour() < 12: f.Println("Bom dia! ", nome)
		case t.Hour() < 18: f.Println("Boa Tarde! ", nome)
		default: f.Println("Boa Noite! ", nome)
	}
}