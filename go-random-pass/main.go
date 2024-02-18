package main

import (
	"fmt"
	"math/rand"
)

func createPassword(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)
	
	for i := 0; i < length; i++ {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

func main() {
	fmt.Println("Insert your length: ")
	var num int
	fmt.Scanln(&num)

	fmt.Println("Your password is: ", createPassword(num))
}