package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // equivalente ao OU exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func imprimir(c1, c2, c3 bool) {
	fmt.Println("ComprarTV50:", c1)
	fmt.Println("ComprarTV32:", c2)
	fmt.Println("ComprarSorvete:", c3)
	fmt.Println("Saudável", !c3) // operador unário
	fmt.Println("--------------")
}

func main() {
	imprimir(compras(true, true))
	imprimir(compras(true, false))
	imprimir(compras(false, false))

}