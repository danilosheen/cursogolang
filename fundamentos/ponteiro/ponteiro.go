package main

import "fmt"

func main() {

	i := 1

	// G0 não possui aritmérica de ponteiros
	var p *int = nil

	p = &i // pegando o endereço da variável i e atribuindo a p

	fmt.Println(*p) // imprime o valor que o ponteiro aponta
	fmt.Println(&i) // pegando o valor na memória de i
	fmt.Println(p) // imprime o valor na memória onde está armazenada o i
	fmt.Println("p e i estão no mesmo lugar da memória?", p == &i)

	*p++ //incrementando a variável i através do ponteiro
	i++ //incrementando a variável i diretamente

	println(*p, i)
}