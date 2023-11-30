package main

import (
	"fmt"
	"math"
)

func drawCircle(radius int) {
	for i := -radius; i <= radius; i++ {
		for j := -radius; j <= radius; j++ {
			distance := math.Sqrt(float64(i*i + j*j))
			if distance < float64(radius) {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func main() {
	// Set the radius of the circle
	radius := 10

	// Call the function to draw the circle
	drawCircle(radius)
}
