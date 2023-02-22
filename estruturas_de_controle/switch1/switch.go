package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota{
	case 10:
		fallthrough

	case 9:
		return "A"

	case 7, 8:
		return "B"

	case 5, 6:
		return "C"

	case 3, 4:
		return "D"

	case 0, 1, 2:
		return "E"
	
	default:
		return "Nota inválida"
	}
	
}

func main()  {
	
	fmt.Println(notaParaConceito(6.2))
	fmt.Println(notaParaConceito(9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(8.9))
}

