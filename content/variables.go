package content

import (
	"fmt"
)

func variables() {

	// Variables

	var printValue string = "Hello World!"
	printMe(printValue)

	var intNum int = 5
	fmt.Println(intNum)

	var floatNum float64 = 1.523
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result1 float32 = floatNum32 + float32(intNum32)
	fmt.Println(result1)

	var intNum1 int = 3
	var intNum2 int = 2

	// arithmetic
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	// string variable and adding a newline
	var myString string = "Hello\nWorld"
	fmt.Println(myString)

	// bool variable and shortening the variable declaration
	myBoolean := false
	fmt.Println(myBoolean)

	var1, var2, var3, var4 := 1, 2, 3, 4

	fmt.Println(var1, var2, var3, var4)

	//var myVar string = foo()
	//fmt.Println(myVar)

	const myConst string = "const value"
	fmt.Println(myConst)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}
