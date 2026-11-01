package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
}

func (Person) Speak() string {
	return "Hola mundo"
}

type MyErr struct {
}

func (e *MyErr) Error() string {
	return "crashed"
}
func returnTypedNil() error {
	var e *MyErr = nil
	return e

}

func realNil() error{
	return nil
}
func main() {
	err1 := returnTypedNil()

	fmt.Println(err1.Error())
}