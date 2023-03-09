package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s %s (interação %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	go fale("Maria", "Estou executando em concorrência", 10)
	fale("Danilo", "Estou executando na principal", 5)
}