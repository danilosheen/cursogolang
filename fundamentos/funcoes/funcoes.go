package main

import "fmt"

func somar(x int, y int) int {
	return x + y
}

func imprimir(valor int) {
	fmt.Println(valor)
}

//go build - irá criar um .exe para testar em caso de modularização