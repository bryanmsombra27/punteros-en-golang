package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Profile struct {
	Bio string
}

type User struct {
	Name    string
	Address Address
	Profile *Profile //Opcional  puede usarse punteros para decir que un valor es opcional o nulo
}

type Counter struct {
	number int
}

func (counter *Counter) Increment(){
	counter.number ++
}

func (counter Counter) Value() int {
return	counter.number	
}



func main() {
	user := User{Name: "ana", Address: Address{City: "CDMX", State: "CDMX"}, }

	if user.Profile == nil {
		fmt.Println("Sin perfil")
	}
	user.Profile = &Profile{
		Bio: "Soy developer mudher fucker",
	}

	fmt.Println(user)
	fmt.Println(user.Profile.Bio)
	
	counter:=Counter{10}
counter.Increment()
counter.Increment()
	fmt.Println(counter.Value())

}