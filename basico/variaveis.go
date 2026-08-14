package main

import (
	f "fmt"
	m "math"
	r "reflect"
)

func main() {	
	const PI = 3.14
	var raio float64

	f.Print("Informe um valor para o raio:")
	f.Scan(&raio)
 
	area := PI *  m.Pow(raio, 2)
	f.Printf("Área é de %.2f\n", area)
}