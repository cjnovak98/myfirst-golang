package main

import ("fmt"
	"golang.org/x/term"
	"os")

func main() {
	var name string
	fmt.Print("Please enter in some stuff for me: ")
	fmt.Scan(&name)
	fmt.Println("You entered: ", name)


	fmt.Println("Now let's do it again, but secretly")
	input string, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println("You entered: ", input)
}
