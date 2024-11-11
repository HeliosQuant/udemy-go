package main

import (
	"fmt"
	"os"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return (t.height * t.base) / 2
}

type square struct {
	side float64
}

func (s square) getArea() float64 {
	return float64(s.side * s.side)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

//////////////////////////////////

func main() {
	t := triangle{height: 10, base: 5}
	sq := square{side: 4}

	printArea(t)
	printArea(sq)

	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	bs := make([]byte, 99999)
	n, err := file.Read(bs)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(bs[:n]))
}
