package main

import "fmt"

func closure() func() {
	x := 10

	var funcao = func() {
		fmt.Println(x)
	}
	return funcao

}

func main() {
	x := 10
	fmt.Println(x)
	imprimeX := closure()
	imprimeX()

}