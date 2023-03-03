package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string){
	fmt.Printf("F2: %s, %s\n", p1, p2)
}

func f3() string{
	return "F3"
}

func f4(v1, v2 float64) float64 {
	return (v1 + v2) / 2
} 

// no go as funções podem ter mais de um retorno
func f5()(string, string){
	return "retorno1", "retorno2"
}

func main(){
	f1()
	f2("String 1", "String 2")
	fmt.Printf("F3: %v\n", f3())
	fmt.Println(f4(10, 5))
	fmt.Println(f5())
	// usando o underline é possível ignorar algum dos valores de retorno
	_, rf52 := f5()
	fmt.Println(rf52)
}	