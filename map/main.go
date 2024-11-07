package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[int]string)
	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"green": "#00ff00",
	// }

	colors[10] = "#FFFFFF"

	fmt.Println(colors)

	delete(colors, 10) // built-in function dedicated for deleteing map element
	fmt.Println(colors)
}
