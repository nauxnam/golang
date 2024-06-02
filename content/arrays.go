package content

import (
	"errors"
	"fmt"
)

func arrays() {
	// ARRAYS
	// Initializing array
	//var intArr [3]int32 = [3]int32{1, 2, 3}
	//fmt.Println(intArr)

	//Slicing, appending array
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	//  |---capacity---|
	// [4, 5, 6, 7, *, *]
	//	|-length-|
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// MAPS

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	// Maps, always returns something even specified var is not in map. So its a good practice
	// to have a bool if else statement.
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)

	} else {
		fmt.Println("Invalid Name")
	}
	// for loop with keys
	for name, age := range myMap2 {
		fmt.Printf("Name: %v\n", name, age)
	}
	//
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// -- -> decrements by 1, ++ -> increments by 1, i += 10 -> increments by 10, i -=10 -> decrements by 10
	//  i *= 10 -> multiply by 10, i /= 10 -> divide by 10
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err

}
