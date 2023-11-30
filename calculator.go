
/*
Package main is the entry point for the executable program.
It contains the main function and other utility functions for a simple calculator.

The calculator supports basic arithmetic operations such as addition, subtraction, multiplication, and division.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getInput function reads user input from the console, converts it to a float64, and returns the result.
func getInput(prompt string, reader *bufio.Reader) (float64, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

// Arithmetic functions for basic operations

// add function performs addition of two numbers.
func add(x, y float64) float64 {
	return x + y
}

// subtract function performs subtraction of two numbers.
func subtract(x, y float64) float64 {
	return x - y
}

// multiply function performs multiplication of two numbers.
func multiply(x, y float64) float64 {
	return x * y
}

// divide function performs division of two numbers.
func divide(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Error: Cannot divide by zero.")
		return 0
	}
	return x / y
}

func main() {
	// Setting up a reader to read input from the console
	reader := bufio.NewReader(os.Stdin)

	// Displaying a welcome message
	fmt.Println("Simple Calculator")

	// Main loop for the calculator
	for {
		// Displaying user options
		fmt.Println("\nOptions:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")

		// Getting user choice
		choice, err := getInput("Enter your choice (1-5): ", reader)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handling user choices
		switch choice {
		case 1, 2, 3, 4:
			// Getting two numbers from the user for the chosen operation
			num1, err := getInput("Enter first number: ", reader)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			num2, err := getInput("Enter second number: ", reader)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			var result float64
			// Performing the selected operation
			switch choice {
			case 1:
				result = add(num1, num2)
			case 2:
				result = subtract(num1, num2)
			case 3:
				result = multiply(num1, num2)
			case 4:
				result = divide(num1, num2)
			}

			// Displaying the result
			fmt.Printf("Result: %0.2f\n", result)

		case 5:
			// Exiting the program if the user chooses option 5
			fmt.Println("Exiting the calculator. Goodbye!")
			return

		default:
			// Handling invalid choices
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}
