package content

import (
	"fmt"
)

// -- -> decrements by 1, ++ -> increments by 1, i += 10 -> increments by 10, i -=10 -> decrements by 10
//  i *= 10 -> multiply by 10, i /= 10 -> divide by 10

func functions() {

	// if checks
	if 1 == 1 && 2 == 2 {
		fmt.Println("Passed the check")
	} //given two checks needs to be true to print line.

	if 1 == 1 || 2 == 2 {
		fmt.Println("Passed the check")
	} //only one check needs to be true to print line. if first check is true, will ignore the rest.

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division %v\n", result)
	default:
		fmt.Printf("\nThe result of the integer division is %v with remainder %v\n", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact\n")
	case 1, 2:
		fmt.Printf("The division was close\n")
	default:
		fmt.Printf("The division was not close\n")
	}
}
