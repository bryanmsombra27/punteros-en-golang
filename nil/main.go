package main

import "fmt"

type User struct {
	Name string
}

type Config struct {
	Host string
	Port int
}

func PrintName(user *User) {
	if user == nil{
		fmt.Println("Usuario indefiinido")
		return 
	}

	fmt.Println("User: ",user.Name)
}
func NewConfig() *Config{
	return &Config{Host: "localhost",Port: 8080}
}


func main() {
	var user *User

	PrintName(user)
	
	user= &User{Name: "Bryan"}
	PrintName(user)

}