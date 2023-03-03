package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(3, 4))

	sub := func(x, y int) int {
		return x - y
	}

	fmt.Println(sub(5, 2))
}