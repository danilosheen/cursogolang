package main

import (
	"fmt"
	"reflect"
)

func main() {

	var b byte = 3
	fmt.Println(b)

	i := 3 //inferência de tipo
	fmt.Println(reflect.TypeOf(i))
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 1 // i = i / 1
	i *= 1 // i = i * 1
	i %= 2 // i = i % 1

	fmt.Println(i)

	//atribuição múltipla
	x, y := 1, 2
	fmt.Println(x, y)

	//troca de valor entre variáveis
	x, y = y, x
	fmt.Println(x, y)

}