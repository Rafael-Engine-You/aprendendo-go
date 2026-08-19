package main

import "fmt"

func calcularMedia (notas []float64) float64 {
	total := 0.0
	contador := 0
	for contador < len(notas) {
		total += notas[contador]
		contador++
	} 
	return total / float64(len(notas))
}

func main() {

	notas := []float64 {7.5,4.0,8.2,6.5,9.0,3.5,5.0,7.0,10.0,6.0}

	qtdAprovados := 0
	qtdReprovados := 0
	qtdRecuperacao := 0

	maiorNota := 0.0
	menorNota := 10.0

	mediaGeral := calcularMedia(notas)

	for i := 0; i < len(notas); i++ {

		if notas[i] >= 7 {
			qtdAprovados++
		} else if notas[i] >= 5 && notas[i] < 7 {
			qtdRecuperacao++
		} else {
			qtdReprovados++
		}

		if maiorNota < notas[i]  {
			maiorNota = notas[i]
		} 
		
		if menorNota > notas[i] {
			menorNota = notas[i]
		}
	}
	
	fmt.Printf("Aprovados: %d | Recuperação: %d | Reprovados: %d\n", qtdAprovados, qtdRecuperacao, qtdReprovados)
	fmt.Printf("Maior nota: %.2f | Menor Nota: %.2f  | Média: %.2f\n", maiorNota, menorNota, mediaGeral)
}