package main

import "fmt"

//channels block
//Bi - Directional channel
//No buufer
//Second c <-43 would break
func main() {

	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	//bufferd

	c2 := make(chan int, 2)
	go func() {
		c2 <- 42
		c2 <- 43
	}()
	fmt.Println(<-c2)
}
