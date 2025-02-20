package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Printf("You can make it! Of the %v mile trip, you've got %v miles left in the tank\n", miles, e.milesLeft()-miles)
	} else {
		fmt.Printf("You can't make it bucko. Fill up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Printf("You've got %v gallons\n\n", myEngine.gallons)
	canMakeIt(myEngine, 100)

}
