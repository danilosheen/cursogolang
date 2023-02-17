package main

import (
	"fmt"
	// usando alias no math
	m "math"
	// _ "math" usando import com underline pode ser importado sem ser usado
 )

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	// forma reduzida de declarar e atribuir uma variável
	area := PI * m.Pow(raio, 2)

	//printf usado para imprimir formatado
	fmt.Printf("A área do círculo é: %.2f \n", area)

	const(
		a = 1
		b = 2
	)

	var( 
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	//a variável f depois de criada como bool, sempre será bool
	var e, f bool = true, false 

	fmt.Println(e, f)

	g, h, i, f := 2, 10.2, "hello", true

	fmt.Println(g, h, i, f)
}