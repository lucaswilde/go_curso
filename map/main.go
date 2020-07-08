package main

import "fmt"

func main() {
	// colors := make(map[int]string)
	// colors[10] = "ffffff"
	// delete(colors, 10)
	// fmt.Println(colors)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#abc123",
		"white": "ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	// color = indice
	// hex = valor
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
