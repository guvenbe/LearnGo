package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		//log.Fatalln("err happened:", err)
		//log.Println("err happened:", err)
		log.Panic("err happened:", err)
	}
}

func foo() {
	fmt.Println("when os.Exit() is called, deferred function don't run")
}
