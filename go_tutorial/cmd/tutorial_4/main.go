package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3}
	intArr[1] = 111
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Eve": 45}

	var age, ok = myMap2["Adam"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v, Address: %v\n", i, v, &i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
