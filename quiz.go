package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Question struct represents a quiz question
type Question struct {
	QuestionText string
	Answer       string
}

// Quiz struct holds a collection of questions and the user's score
type Quiz struct {
	Questions []Question
	Score     int
}

// NewQuiz creates a new quiz with a set of questions
func NewQuiz() Quiz {
	questions := []Question{
		{"What is the capital of France?", "Paris"},
		{"Which planet is known as the Red Planet?", "Mars"},
		{"What is the largest mammal?", "Blue Whale"},
		{"Who wrote 'Romeo and Juliet'?", "William Shakespeare"},
		{"What is the currency of Japan?", "Yen"},
	}

	return Quiz{Questions: questions, Score: 0}
}

// RunQuiz presents the quiz questions and collects user answers
func (q *Quiz) RunQuiz() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Quiz!")

	// Ask 5 questions
	for i := 0; i < 5; i++ {
		question := q.Questions[i]
		fmt.Println(question.QuestionText)
		userAnswer, _ := getUserInput("Your answer: ", reader)

		// Check if the answer is correct
		if strings.EqualFold(userAnswer, question.Answer) {
			fmt.Println("Correct!")
            fmt.Println()
			// Increment the score
			q.Score += 20
		} else {
			fmt.Printf("Incorrect. The correct answer is: %s\n\n", question.Answer)
		}
	}

	// Print the total score out of 100
	fmt.Printf("Quiz completed. Your total score: %d out of 100\n", q.Score)
}

// getUserInput gets user input from the console
func getUserInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func main() {
	quiz := NewQuiz()
	quiz.RunQuiz()
}
