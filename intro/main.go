package main

import "fmt"

func main() {
	x := 10
	p := &x //acceder a la direccion de memoria  mediante el & 

	fmt.Println("x= ",x)
	fmt.Println("p= ",p)//direccion de memoria: 0x78b9b60e060
	fmt.Println("*p= ",*p) // al agregarle un * se accede al valor de la direccion de memoria

	//con el  * se lee/escribe el valor de la direccion en memoria
	*p = 99

	fmt.Println("x= ",x)// el valor tambien es modificado ya que apunta a la misma direccion de memoria
	fmt.Println("p= ",p)//direccion de memoria: 0x78b9b60e060
	fmt.Println("*p= ",*p) // al agregarle un * se accede al valor de la direccion de memoria
}
