package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- foo()
	}()
	go func() {
		ch2 <- bar()
	}()
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func foo() int {
	var sum int
	for i := 0; i < 100000; i++ {
		fmt.Println("foo", i)
		sum += i
	}
	return sum
}
func bar() int {
	var sum int
	for i := 0; i < 100000; i++ {
		fmt.Println("bar", i)
		sum += i
	}
	return sum
}
