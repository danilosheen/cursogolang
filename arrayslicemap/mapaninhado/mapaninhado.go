package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"C": {
			"Cicero Danilo":  5000.65,
			"Cicero Pereira": 3600.89,
			"Cicero Silva":   6000.25,
		},
		"D": {
			"Daniel José": 9000.54,
		},
		"M": {
			"Maria Claysla": 5000.65,
		},
	}

	delete(funcsPorLetra, "D")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs{
			fmt.Println(nome, salario)
		}
	}
}
