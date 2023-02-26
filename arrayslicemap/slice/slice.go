package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]

	fmt.Println(a2, s2)

	/*
	- o slice é uma estrutura que aponta para o primeiro valor referenciado 
	de um array existente

	- slice não é considerado um array

	- é possível criar vários slices apontando para o mesmo array em posições
	diferentes
	*/

	s3 := a2[:2]
	fmt.Println(s3)
}