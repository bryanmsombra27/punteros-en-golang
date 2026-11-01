package main

import "fmt"

type User struct {
	Name     string
	Age      int
	IsActive bool
}

type UpdateUser struct {
	Name     *string
	Age      *int
	IsActive *bool
}

func Apply(user *User,update UpdateUser){
	if update.Name != nil {
		user.Name = *update.Name
	}
	if update.Age != nil {
		user.Age = *update.Age
	}
	if update.IsActive != nil {
		user.IsActive = *update.IsActive
	}

}



func main() {

	user := User{Name: "ana", Age: 20, IsActive: true}

	fmt.Println(user)
	
	newName:= "Rodrigo"
	zeroAge:=0
	isActive:= false
	
	Apply(&user,UpdateUser{Name: &newName,Age: &zeroAge,IsActive: &isActive})
	fmt.Println(user)
	
}