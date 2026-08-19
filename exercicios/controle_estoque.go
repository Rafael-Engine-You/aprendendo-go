package main

import "fmt"

func main() {

	semEstoque := 0
	critico := 0
	baixo := 0
	normal := 0
	maiorEstoque := 0

	produtoIndice := 0
	totalEstoque := 0

	estoque := []int{10, 0, 3, 25, 7, 1, 0, 15}

	for i := 0; i < len(estoque); i++ {
		
		if estoque[i] >= 1 && estoque[i] <= 3  {
			critico++  
		} else if estoque[i] >= 4 && estoque[i] <= 10 {
			baixo++
		} else if estoque[i] > 10 {
			normal++
		} else {
			semEstoque++
		}

		if estoque[i] > maiorEstoque {
			maiorEstoque = estoque[i]
			produtoIndice = i + 1
			totalEstoque = estoque[i]
		}
	}

	fmt.Printf("Sem estoque: %d | Critico: %d | Baixo: %d | Normal: %d\n", semEstoque, critico, baixo, normal)
	fmt.Printf("Maior estoque: Produto %d (indice %d), com %d unidades\n", produtoIndice, produtoIndice - 1, totalEstoque)
}