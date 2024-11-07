package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//alex := person{firstName: "helios", lastName: "quant"}
	var alex person
	alex.firstName = "helios"
	alex.lastName = "quant"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
