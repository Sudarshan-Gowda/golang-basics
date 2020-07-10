package main

import "fmt"

func main() {

	var a int
	var b int
	var operation string
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter the second  number: ")
	fmt.Scanln(&b)
	fmt.Print("Enter the operation to perform: ")
	fmt.Scanln(&operation)

	if operation == "add" {
		fmt.Println("addition of 2 number is: ", add(a, b))
	} else if operation == "sub" {
		fmt.Println("substraction of 2 number is: ", sub(a, b))
	} else if operation == "mul" {
		fmt.Println("multiplication of 2 number is: ", mul(a, b))
	} else if operation == "div" {
		fmt.Println("duvision of 2 number is: ", div(a, b))
	} else {
		fmt.Println("Invalid Operation!!")
	}
}
