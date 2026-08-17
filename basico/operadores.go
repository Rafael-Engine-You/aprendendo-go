package main

import (
	f "fmt"
	r "reflect"
)

func main() {
	
	var b byte = 3 
	f.Println("Tipo variavel b ", r.TypeOf(b))

	i := 3 // inferencia de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2  
	f.Println("i", resultado)

	x, y := 1,2
	f.Println(x,y)

	// operadores unarios
	x = x + 1 //x++ x--

}