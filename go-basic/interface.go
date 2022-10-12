package main

import "fmt"

type HasName interface{
	GetName() string
}

func SayHello(hasName HasName){
	fmt.Println("Hello", hasName.GetName())
}

type Person struct{
	Name string
}

func (person Person) GetName() string{
	return person.Name
}

func main() {
	var rzq Person
	rzq.Name = "Rzq"
	SayHello(rzq)
}