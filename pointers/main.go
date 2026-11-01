package main

import "fmt"



func increment(number *int) {
	*number++
}

type User struct {
	Name string
	Age int
}
func birthDay(user *User){

	user.Age++
}

func main() {
	x := 10

	fmt.Println(x)
	increment(&x)
	fmt.Println(x)

	user:= User{
		Name: "albedo",
		Age: 12,
	}

	birthDay(&user)
	 a,b:=1,2


	fmt.Println(a,b)
	swap(&a,&b)
	fmt.Println(a,b,"SWAP")
	

}

func swap(a *int, b *int){
	c:=*b

	*b=*a
	*a=c

  


}