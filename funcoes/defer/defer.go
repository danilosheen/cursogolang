package main

import "fmt"

func obterValorAprovado(valor int) int {
	defer fmt.Println("Fim, antes do return")

	if valor > 5000 {
		fmt.Println("Valor alto!!")
		return valor
	}

	fmt.Println("Valor baixo")
	return valor

}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(4000))
}
