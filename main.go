package main

import (
	"fmt"
)

func main() {

	// var colors map[string]string // a way to create an empty map

	// colors := make(map[string]string) // another way to create an empty map
	// colors["white"] = "#ff0000"       // then adding the value to that empty map

	// colors := make(map[int]string) // map with key as int & value as string
	// colors[10] = "#ff0000"

	// delete(colors, 10) // built in function to delete a value from map

	colors := map[string]string{ // key & value are of type string
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fff745",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for hex, color := range c {
		fmt.Println("hex code for", color, "is", hex)
	}

}
