package main

import "fmt"

func printGrade(score float64) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func main() {
	// Print statement
	fmt.Println("Hello world")

	// User input
	var userName string
	var userScore float64
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)
	fmt.Print("Your name is ", userName)

	fmt.Print("\nEnter your score: ")
	fmt.Scan(&userScore)
	fmt.Print("With a score of ", userScore, " you earned a grade of ", printGrade(userScore))

	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
