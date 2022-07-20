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

	cs := make(chan<- int) //send channel
	cr := make(<-chan int) //receive channel
	fmt.Printf("")
	fmt.Printf("---------------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("---------------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	//send
	go foo(c)
	//receive
	bar(c)
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
