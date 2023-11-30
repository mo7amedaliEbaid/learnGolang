package main

import "fmt"

func drawOppositeTriangles(rows int) {
	for i := 1; i <= rows; i++ {
		// Print spaces before the letters
		for j := 1; j <= rows-i; j++ {
			fmt.Print("  ")
		}

		// Print ascending letters
		for j := 0; j < i; j++ {
			fmt.Printf("%c ", 'A'+j)
		}

		// Print descending letters
		for j := i - 2; j >= 0; j-- {
			fmt.Printf("%c ", 'A'+j)
		}

		// Move to the next line
		fmt.Println()
	}
}

/* func main() {
	// Set the number of rows for the triangles
	rows := 5

	// Call the function to draw the opposite triangles with letters
	drawOppositeTriangles(rows)
} */
