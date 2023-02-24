package main

import "fmt"

func main() {
	//estrutura homogênea (mesmo tipo) e estática (tamanho fixo)
	var notas [3]float64

	//atribui valores nas posições do array
	notas[0], notas[1], notas[2] = 3.6, 5.8, 9.9
	fmt.Println(notas)

	//faz o somatório usando o for para percorrer o array
	sum := 0.0
	for i:=0; i < len(notas); i++{
		sum += notas[i]
	}

	//converte o tamanho do array para float64, go não aceita divisão de tipos
	//diferentes
	media := sum / float64(len(notas))
	fmt.Printf("%.2f\n", media)

}