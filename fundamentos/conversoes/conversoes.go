package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 3

	//em casos de divisão usando variáveis de tipagem diferente como float64
	//e int, é necessário fazer a conversão
	fmt.Println(x / float64(y))

	//no caso da conversão direta a casa decimal é ignorada resoltado = 6
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado em caso de conversão de string ele é apresentado valor asc
	fmt.Println("Teste " + string(rune(97)))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(97))

	//string para int
	//underline passa para o go que aquele valor será ignorado
	num, _ := strconv.Atoi("97")
	fmt.Println(num - 96)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}