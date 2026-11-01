package main

import "fmt"

func incrementValue(number int ) {
	number++
}

func mutateFirts(slice []int){

	slice[0] = 10
}

// rompiendo la referencia porque los slice ya no manejan la misma capacidad al diferir en de capacida en automatico se crea una nueva referencia en memoria y por tanto ya no se modifican de la misma forma
func push(slice []int) []int {
	slice = append(slice,42)
	return slice
} 

func main() {
	x := 10
	a:=[]int{1,2,3}

	// paso por valor: es decir se envia una copia del valor original de x
	incrementValue(x)

	// paso por referencia: es decir que se envia el valor original de a. los slice o maps se pasan por referencia siempre
	mutateFirts(a)
	


	fmt.Println(x)
	fmt.Println(a)
	
	b:= push(a)
	fmt.Println(a,"comparativa a")
	fmt.Println(b, "comparativa b")
}