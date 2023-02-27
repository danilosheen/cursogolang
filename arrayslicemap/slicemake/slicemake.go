package main

import "fmt"

func main() {

	//atribui valor na posição do array apontado pelo slice
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	//adiciona uma capacidade para o array apontado pelo slice
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	//adiciona valores no array apontado pelo slice
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s)

	//mesmo se o array estiver lotado e for adicionado mais um valor, a
	//capacidade é aumentada automaticamente
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}