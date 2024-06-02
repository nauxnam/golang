package content

import (
	"fmt"
)

// STRUCTS

// Structs defined by "type"
type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

//Interfaces

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("You need to fuel up first!")
	}
}

func structs() {
	//STRUCTS

	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 25} // or you can simply omit just values without names.
	// or you can define it by myEngine.mpg = 20 for mpg value
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	var myEleEngine electricEngine = electricEngine{25, 15}
	canMakeIt(myEleEngine, 50)
}
