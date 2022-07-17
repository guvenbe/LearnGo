package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("OS\t", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Add(2)
	go randNumber(1)
	go randNumber(2)
	wg.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}

func randNumber(num int) {
	for i := 0; i < 10000; i++ {
		min := 1
		max := 70
		fmt.Printf("Gen %d - %d\n", num, rand.Intn(max-min+1)+min)

	}
	wg.Done()
}
