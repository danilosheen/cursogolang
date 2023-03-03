package main

import "fmt"

func media(valores ...float64) float64 {
	total := 0.0
	for _, valor := range valores {
		total += valor
	}
	if total == 0.0{
		return 0
	}

	return total / float64(len(valores))
}

func main() {
	media := media(1, 4, 5, 4.3, 10, 9.2, 9.4)
	fmt.Printf("Média: %.2f", media)
}
