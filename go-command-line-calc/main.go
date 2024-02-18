package main

import "fmt"

func main() {
	fmt.Print("Give me first number: ")
	var float1 float64
	fmt.Scanln(&float1)
	
	fmt.Print("Give me second number: ")
	var float2 float64
	fmt.Scanln(&float2)
	
	fmt.Print("Give me an operand: ")
	var operand string
	fmt.Scanln(&operand)
	
	result := 0.0
	
	switch operand {
		case "+":
			result = float1 + float2
		case "-":
			result = float1 - float2
		case "*":
			result = float1 * float2
		case "/":
			result = float1 / float2
		default:
			fmt.Println("Invalid operand")
			return
	}
	
	fmt.Println("Result:", result)
}