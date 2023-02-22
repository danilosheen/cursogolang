package main

import (
	"fmt"
	"time"
)

func main(){

	i := 1
	for i <= 10{ // go não possui while
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println("") // quebra de linha na última linha

	for i := 0; i < 10; i+= 2{
		fmt.Printf("%d ", i)
	}
	fmt.Println("") 

	for i := 1; i <= 10; i++{
		if i%2 == 0{
			fmt.Printf("O número %d é par\n", i)
		} else {
			fmt.Printf("O número %d é ímpar\n", i)
		}
	}

	count := 0
	for { // laço semelhante ao while com pausa em condição
		if count == 5{
			break
		}
		fmt.Println("Imprimindo...")
		time.Sleep(time.Second)
		count++
	}
	for { // laço infinito
		fmt.Println("Imprimindo segundo for...")
		time.Sleep(time.Second)
	}

}