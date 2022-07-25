package main

import (
	"LearnGo/demo/greetings"
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
