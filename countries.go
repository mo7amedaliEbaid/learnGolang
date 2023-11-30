package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CountryInfo represents information about a country
type CountryInfo struct {
	FlagColors []string
	Continent  string
}

func main() {
	countries := map[string]CountryInfo{
		"france": {
			FlagColors: []string{"Blue", "White", "Red"},
			Continent:  "Europe",
		},
		"japan": {
			FlagColors: []string{"Red", "White"},
			Continent:  "Asia",
		},
		"brazil": {
			FlagColors: []string{"Green", "Yellow", "Blue", "White"},
			Continent:  "South America",
		},
		// Add more countries as needed
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Country Info App!")

	for {
		countryName := getUserInput("Enter a country name (or 'exit' to quit): ", reader)

		// Check if the user wants to exit
		if strings.ToLower(countryName) == "exit" {
			fmt.Println("Exiting the Country Info App. Goodbye!")
			return
		}

		// Lookup country information
		info, found := countries[strings.ToLower(countryName)]
		if !found {
			fmt.Println("Country not found. Please enter a valid country name.")
			continue
		}

		// Print country information
		fmt.Printf("Flag Colors: %s\n", strings.Join(info.FlagColors, ", "))
		fmt.Printf("Continent: %s\n", info.Continent)
	}
}

// getUserInput gets user input from the console
func getUserInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input))
}
