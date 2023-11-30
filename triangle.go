package main

import "fmt"

func drawTriangle(rows int) {
	for i := 1; i <= rows; i++ {
		// Print spaces before the numbers
		for j := 1; j <= rows-i; j++ {
			fmt.Print("  ")
		}

		// Print ascending numbers
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}

		// Print descending numbers
		for j := i - 1; j >= 1; j-- {
			fmt.Printf("%d ", j)
		}

		// Move to the next line
		fmt.Println()
	}
}

func main() {
	// Set the number of rows for the triangle
	rows := 5

	// Call the function to draw the triangle
	drawTriangle(rows)
}
