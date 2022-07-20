package main

import "fmt"

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	//send
	go send(eve, odd, quit)
	//receive
	receive(eve, odd, quit)
	fmt.Println("about to exit")
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("from the eve channel", v)
		case v := <-odd:
			fmt.Println("from the odd channel", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma not ok", i, ok)
				return
			} else {
				fmt.Println("from comma ok", i)
			}
		}
	}
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)

}
