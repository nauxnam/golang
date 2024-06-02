package content

import "fmt"

func channels() {
	var c = make(chan int)
	c <- 1
	var i = <-c
	fmt.Println(i)
}
