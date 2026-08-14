package main

import (
	f "fmt"
	r "reflect"
)

func main() {
	f.Println("Tipo de variáveis")

	nome := "Maria José"
	f.Println("Tipo variável nome", r.TypeOf(nome))
	idade := 24
	f.Println("Tipo variável idade", r.TypeOf(idade))
	eMaior := true
	f.Println("Tipo variável eMaior", r.TypeOf(eMaior))
	peso := 45.9
	f.Println("Tipo variável peso", r.TypeOf(peso))
}