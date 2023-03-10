package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Executou")
	ch <- 4
	ch <- 5
	ch <- 6
}

//nesse exemplo o buffer do canal tem tamanho 3, isso faz com que após o canal
//ser preenchido ele fica bloqueado para novos valores, só conseguindo alocar
//mais valores após ser consumido

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}