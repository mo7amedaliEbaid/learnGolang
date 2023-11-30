package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Translator represents a simple translation app
type Translator struct {
	Translations map[string]string
}

// NewTranslator initializes a new Translator with some default translations
func NewTranslator() Translator {
	return Translator{
		Translations: map[string]string{
			"hello": "bonjour",
			"world": "monde",
			"go":    "aller",
			// Add more translations as needed
		},
	}
}

// Translate translates a word to another language
func (t *Translator) Translate(word string) string {
	translation, found := t.Translations[strings.ToLower(word)]
	if !found {
		return "Translation not found"
	}
	return translation
}

func main() {
	translator := NewTranslator()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Translator App!")

	for {
		word := getUserInput("Enter a word to translate (or 'exit' to quit): ", reader)

		// Check if the user wants to exit
		if strings.ToLower(word) == "exit" {
			fmt.Println("Exiting the Translator App. Goodbye!")
			return
		}

		// Translate and print the result
		result := translator.Translate(word)
		fmt.Printf("Translation: %s\n", result)
	}
}

// getUserInput gets user input from the console
func getUserInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input))
}
