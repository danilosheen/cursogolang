package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // nesse caso vai haver retorno limpo pois ja foi informado o que a função vai retornar
}

func main() {
	fmt.Println(trocar(1, 2))
}