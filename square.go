package main

import "fmt"

func drawSquare(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
/*
func main() {
	// Set the size of the square
	size := 5

	// Call the function to draw the square
	drawSquare(size)
}
*/