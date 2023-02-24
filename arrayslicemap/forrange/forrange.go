package main

import "fmt"

func main() {

	numeros := [...]int{1, 2, 3, 4, 5, 6} //compilador faz a contagem do array

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros{
		fmt.Println(num)
	}
}