package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ffOOOO",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["yellow"] = "#000000"
	delete(colors, "yellow")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
