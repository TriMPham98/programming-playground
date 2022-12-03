package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var userName string
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)
	fmt.Print("Your name is ", userName)
}
