package main

import "fmt"

var (
	name, location string
	age            int
)

// or one by

var name1 string
var age1 int
var location1 string

//initialize

var (
	name3     string = "Prince Oberyn"
	age3      int    = 32
	location3 string = "Dorne"
)

//If an initializer is present, the type can be omitted
var (
	name4     = "Prince Oberyn"
	age4      = 32
	location4 = "Dorne"
)

//You can also initialize multiple variables at once.
var (
	name5, location5, age5 = "Prince Oberyn", "Dorne", 32
)

func main() {
	name, location := "Prince Oberyn", "Dorne"
	age := 32
	fmt.Printf("%s age %d from %s \n", name, age, location)
	fmt.Printf("%T\n", age)  //type
	fmt.Printf("%b\n", age)  //binary
	fmt.Printf("%x\n", age)  //hex
	fmt.Printf("%#x\n", age) //hex with 0x in front
	fmt.Printf("%v", age)    //value in default format

}
