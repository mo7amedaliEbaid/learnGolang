package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Error: Cannot divide by zero.")
		return 0
	}
	return x / y
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Calculator")

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")

		choice, err := getInput("Enter your choice (1-5): ", reader)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch choice {
		case 1, 2, 3, 4:
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

			fmt.Printf("Result: %0.2f\n", result)

		case 5:
			fmt.Println("Exiting the calculator. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}
