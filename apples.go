/*
Given a total of 100 apples,
distribute them among 10 children of varying ages.
Each child should receive two more apples than the child immediately younger.
Please provide the number of apples allocated to each child.
*/

package main

import "fmt"

func main() {
	// Total number of apples
	totalApples := 100

	// Number of children
	numChildren := 10

	// Calculate the number of apples each child gets
	// Each child gets two more apples than the younger child
	applesPerChild := totalApples / numChildren
	extraApples := 2

	// Display the distribution of apples
	for i := 1; i <= numChildren; i++ {
		childApples := applesPerChild + (i-1)*extraApples
		fmt.Println(childApples)
	}
}
