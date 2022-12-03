package main

import "fmt"

func main() {
	// Print statement
	fmt.Println("Hello world")

	// User input
	var userName string
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)
	fmt.Print("Your name is ", userName)
}
