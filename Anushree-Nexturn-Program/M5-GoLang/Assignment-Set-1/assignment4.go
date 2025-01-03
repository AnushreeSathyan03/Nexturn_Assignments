package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Question struct represents a quiz question
type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

// Questions bank
var questions = []Question{
	{"What is the size of an int in Go on a 64-bit architecture?", [4]string{"32 bits", "64 bits", "Depends on the compiler", "8 bytes"}, 2},
	{"Which keyword is used to declare a new type in Go?", [4]string{"type", "struct", "class", "declare"}, 1},
	{"How do you declare a map in Go?", [4]string{"map[string]int", "map(int, string)", "map<string, int>", "dict[string]int"}, 1},
	{"Which operator is used for pointer dereferencing in Go?", [4]string{"&", "*", "->", "@"}, 2},
	{"What is the purpose of the 'select' statement in Go?", [4]string{"To switch between channels", "To iterate over arrays", "To handle errors", "To terminate a goroutine"}, 1},
	{"What is the output of 'len([]int{1, 2, 3})'?", [4]string{"0", "3", "2", "undefined"}, 2},
	{"Which function is used to format strings in Go?", [4]string{"format()", "strings.Join()", "fmt.Sprintf()", "fmt.Fprintln()"}, 3},
	{"What is a goroutine in Go?", [4]string{"A function running in its own thread", "A lightweight thread managed by Go", "An error handling mechanism", "A type of Go data structure"}, 2},
	{"How do you initialize a slice in Go?", [4]string{"[]int{1, 2}", "slice<int>{1, 2}", "int[] a = {1, 2}", "new(int[1, 2])"}, 1},
	{"What does the 'cap' function return in Go?", [4]string{"The capacity of a slice", "The memory address of a variable", "The length of a string", "The type of an array"}, 1},
}

const questionTimeLimit = 15 * time.Second

func main() {
	fmt.Println("Welcome to the Enhanced Quiz System!")
	fmt.Println("Instructions:")
	fmt.Println("- Enter the option number to select an answer.")
	fmt.Println("- Type 'exit' to quit the quiz anytime.")
	fmt.Println("- You have 15 seconds to answer each question.")

	rand.Seed(time.Now().UnixNano())
	takeQuiz()
}

// takeQuiz handles the quiz process
func takeQuiz() {
	reader := bufio.NewReader(os.Stdin)
	shuffledQuestions := shuffleQuestions(questions)
	score := 0
	totalQuestions := len(shuffledQuestions)

	for i, question := range shuffledQuestions {
		fmt.Printf("Question %d: %s\n", i+1, question.Question)
		for idx, option := range question.Options {
			fmt.Printf("%d. %s\n", idx+1, option)
		}

		answerChan := make(chan string)
		go func() {
			fmt.Print("Your answer: ")
			input, _ := reader.ReadString('\n')
			answerChan <- strings.TrimSpace(input)
		}()

		select {
		case answer := <-answerChan:
			if strings.ToLower(answer) == "exit" {
				fmt.Println("Exiting the quiz...")
				return
			}

			choice, err := strconv.Atoi(answer)
			if err != nil || choice < 1 || choice > 4 {
				fmt.Println("Invalid input! Please enter a valid option (1-4).")
				continue
			}

			if choice == question.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Printf("Incorrect! The correct answer was: %s\n", question.Options[question.Answer-1])
			}

		case <-time.After(questionTimeLimit):
			fmt.Printf("\nTime's up! The correct answer was: %s\n", question.Options[question.Answer-1])
		}

		fmt.Println()
	}

	calculateScore(score, totalQuestions)
}

// calculateScore calculates and displays the user's score and performance
func calculateScore(score, total int) {
	fmt.Printf("\nQuiz Completed! Your score: %d/%d\n", score, total)
	percentage := (float64(score) / float64(total)) * 100

	fmt.Printf("Percentage: %.2f%%\n", percentage)
	if percentage >= 90 {
		fmt.Println("Performance: Outstanding!")
	} else if percentage >= 75 {
		fmt.Println("Performance: Very Good!")
	} else if percentage >= 50 {
		fmt.Println("Performance: Good.")
	} else {
		fmt.Println("Performance: Needs Improvement.")
	}
}

// shuffleQuestions randomizes the order of the questions
func shuffleQuestions(qs []Question) []Question {
	shuffled := make([]Question, len(qs))
	copy(shuffled, qs)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}