package content

import (
	"fmt"
)

func pointers() {
	//Initialize the pointer
	var p *int32 = new(int32)
	var i int32
	//Reference the value of pointer
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("\nThe value if i is %v\n", i)
	*p = 10
	p = &i
	fmt.Printf("\nThe value p points to is: %v\n", *p)
	fmt.Printf("\nThe value if i is %v\n", i)
	//Defining a variable to replace pointers value to defined variable's value.
	var k int32 = 2
	i = k

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result2 [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v\n", result2)
	fmt.Printf("\nThe value of thing1 is: %v\n", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
