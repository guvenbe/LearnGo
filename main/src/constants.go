package main

import "fmt"

const Pi = 3.14
const (
	StatusOK                   = 200
	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent            = 204
	StatusResetContent         = 205
	StatusPartialContent       = 206
)

const (
	Pi2   = 3.14
	Truth = false
	Big   = 1 << 62
	Small = Big >> 61
)

func main() {
	const Greeting = "ハローワールド" //declaring a constant
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
}
