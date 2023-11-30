package main

import (
	//"fmt"
	//"strconv"
)

// Quran struct represents the Quran with surahs and their order
type Quran struct {
	Surahs map[string]int
}

// NewQuran initializes a new Quran with surahs and their order
func NewQuran() Quran {
	surahs := map[string]int{
		"Al-Fatiha (The Opening)":         1,
		"Al-Baqarah (The Cow)":            2,
		"Al-Imran (The Family of Imran)":  3,
		"An-Nisa' (The Women)":            4,
		"Al-Ma'idah (The Table Spread)":   5,
		"Al-An'am (The Cattle)":           6,
		"Al-A'raf (The Heights)":          7,
		"Al-Anfal (The Spoils of War)":   8,
		"At-Tawbah (The Repentance)":     9,
		"Yunus (Jonah)":                  10,
		"Hud (Hud)":                      11,
		"Yusuf (Joseph)":                 12,
		"Ar-Ra'd (The Thunder)":          13,
		"Ibrahim (Abraham)":              14,
		"Al-Hijr (The Rocky Tract)":      15,
		"An-Nahl (The Bee)":              16,
		"Al-Isra' (The Night Journey)":   17,
		"Al-Kahf (The Cave)":             18,
		"Maryam (Mary)":                  19,
		"Taha":                           20,
	}

	return Quran{Surahs: surahs}
}

/* func main() {
	quran := NewQuran()

	// Print the number of surahs
	fmt.Printf("The Quran contains %d surahs.\n", len(quran.Surahs))

	// Print each surah and its order
	for surah, order := range quran.Surahs {
		fmt.Printf("Order: %d, Surah: %s\n", order, surah)
	}

	// Prompt the user to enter a surah order
	var input string
	fmt.Print("Enter the order of a surah: ")
	fmt.Scanln(&input)

	// Convert user input to integer
	userOrder, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid surah order.")
		return
	}

	// Find and print the surah name based on the user's input
	for surah, order := range quran.Surahs {
		if order == userOrder {
			fmt.Printf("Surah with order %d is: %s\n", order, surah)
			return
		}
	}

	fmt.Println("Surah not found for the given order.")
} */
