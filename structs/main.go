package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "helios", lastName: "quant"}
	fmt.Println(alex)
}
