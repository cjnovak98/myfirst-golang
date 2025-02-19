package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	var name string
	fmt.Print("Please enter in some stuff for me: ")
	fmt.Scan(&name)
	fmt.Println("You entered: ", name)

	fmt.Print("Now let's do it again, but secretly: ")
	inputBytes, _ := term.ReadPassword(int(os.Stdin.Fd()))
	input := string(inputBytes)
	fmt.Println("You entered: ", input)
}
