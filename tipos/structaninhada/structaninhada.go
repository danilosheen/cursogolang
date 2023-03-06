package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, qtde: 2, preco: 10.10},
			{produtoID: 2, qtde: 4, preco: 12.10},
			{produtoID: 3, qtde: 5, preco: 16.10},
	
		},
	}

	fmt.Printf("O valor total do pedido é %.2f", pedido.valorTotal())
}
