package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func readUserInputSecretly() string {

	fmt.Print("Now let's do it again, but secretly: ")
	inputBytes, _ := term.ReadPassword(int(os.Stdin.Fd()))
	input := string(inputBytes)
	fmt.Println()

	return input

}

func readUserInput() string {

	var input string
	fmt.Print("Please enter in some stuff for me: ")
	fmt.Scan(&input)

	return input

}

func main() {

	var result string = readUserInput()
	fmt.Println("You entered: ", result)

	var secureresult string = readUserInputSecretly()
	fmt.Println("You securely entered: ", secureresult)
}
