package main

import "fmt"

func multiplica(v1, v2 float64) float64 {
	return v1 * v2
}

func exec(funcao func(float64, float64) float64, v1, v2 float64) float64 {
	return funcao(v1, v2)
}

func main() {

	resultado := exec(multiplica, 3, 4)
	fmt.Println(resultado)
}