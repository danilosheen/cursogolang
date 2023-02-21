package main

import "fmt"

// Go não possui operador ternário
func obterResultado(nota float64) string {
	if nota >= 6.0 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {	
	fmt.Println(obterResultado(6.4))
}