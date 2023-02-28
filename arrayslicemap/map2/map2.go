package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José João":      10000.50,
		"Fábio Ferreira": 11000.25,
		"Regis Brasil":   15000.68,
	}

	funcESalarios["Cicero Danilo"] = 5000.98
	//caso tente remover uma chave inexistente, o go não vai acusar erro
	delete(funcESalarios, "Nenhum") 

	for nome, salario := range funcESalarios {
		fmt.Printf("Nome: %s, Salário: %.2f\n", nome, salario)
	}

}