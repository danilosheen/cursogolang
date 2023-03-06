package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	pessoa1 := pessoa{nome: "Danilo", sobrenome: "Ferreira"}
	fmt.Println(pessoa1)
	fmt.Println(pessoa1.getNomeCompleto())
	pessoa1.setNomeCompleto("Maria Nascimento")
	fmt.Println(pessoa1.getNomeCompleto())

}
