package main

import "fmt"

func main() {
	var intNum uint16 = 1
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum64 float64 = 20.1
	fmt.Println(floatNum64, "This is just printing the float64")
	var intNum32 int32 = 10
	var result float64 = floatNum64 + float64(intNum32)
	fmt.Println(result, "This is adding a float from 20.1 with an int from 10")

	var intNum1 int16 = 3
	var intNum2 int16 = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString, "<<<< This is a string that adds strings together")
	var myString2 string = `Hello
	World`
	fmt.Println(myString2, "<<<< This is a string that formats using back quotes")

	fmt.Println(len("test"), "<<< This is a length of the string 'test'")

	var myBoolean bool = false
	fmt.Println(myBoolean)

	const myConst string = "You can't change this"
	fmt.Println(myConst, "<<<< This is a constant, and is immutable")

}
