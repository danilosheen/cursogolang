package main

import "fmt"


func notaParaConceito(nota float64) string{

	switch{
	case nota >= 9:
		return "A"

	case nota >= 7 && nota <= 8:
		return "B"

	case nota >= 5 && nota < 7:
		return "C"

	case nota >= 3 && nota < 5:
		return "D"

	default:
		return "E"
	}
}

func main()  {
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(9))
	fmt.Println(notaParaConceito(1.2))
}