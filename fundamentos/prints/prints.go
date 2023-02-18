package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := math.Pi

	//como o valor de x é um float, o go não faz concatenação com string
	// fmt.Println("O valor de x é " + x)
	//uma alternativa é usar uma vírgula para exibir o valor de x como floar após
	fmt.Println("O valor de x é", x)

	//outra alternativa é criar uma nova variável e transformar o x em uma string
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	/*
		o printf é a maneira de imprimir com formatações, um exemplo é definir casas
		decimais
	*/

	fmt.Printf("O valor de x é %.2f", x)

	/*
		%f - tipo float
		%s - tipo string
		%d - tipo inteiro
		%t - tipo boolean
		%v - é um wildcard e imprime corretamente diversos formatos
	*/

	a := 1
	b := 1.95
	c := true
	d := "string"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v %v", a, b, b, c, d)

}
