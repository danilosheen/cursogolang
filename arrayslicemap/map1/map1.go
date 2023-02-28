package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//maps devem ser inicializados, caso contrário -> erro map não inicializado

	aprovados := make(map[int]string)

	aprovados[11111111111] = "Pessoa1"
	aprovados[22222222222] = "Pessoa2"
	aprovados[33333333333] = "Pessoa3"
	aprovados[44444444444] = "Pessoa4"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF): %d\n", nome, cpf)
	}

	fmt.Println("")

	fmt.Println(aprovados[22222222222])
	delete(aprovados, 22222222222)

	fmt.Println("")

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF): %d\n", nome, cpf)
	}


}