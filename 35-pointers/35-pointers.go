package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address when you have the address
	*b = 43
	fmt.Println(a)

	x := 0

	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x before", &x)
	fmt.Println("x before", x)

	p1 := person{
		first: "bob",
		last:  "smith",
		age:   55,
	}
	p2 := person{
		first: "steve",
		last:  "jones",
		age:   57,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("--------------")
	changeMe(&p1, &p2)
	fmt.Println(p1)
	fmt.Println(p2)
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}

type person struct {
	first, last string
	age         int
}

func changeMe(p1 *person, p2 *person) {
	(*p1).first = (*p2).first
	p1.last = p2.last
	p1.age = p2.age
}

/*
IMPORTANT: “The method set of a type determines the INTERFACES
that the type implements.....”
~ golang spec
Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 56
The above “important” was left out of this video but will be discussed in the “concurrency”
section in a video called “method sets revisited”.
● a NON-POINTER RECEIVER
○ works with values that are POINTERS or NON-POINTERS.
● a POINTER RECEIVER
○ only works with values that are POINTERS.
Receivers Values
-----------------------------------------------
(t T) T and *T
(t *T) *T
code:
● NON-POINTER RECEIVER & NON-POINTER VALUE
○ https://play.golang.org/p/2ZU0QX12a8
● NON-POINTER RECEIVER & POINTER VALUE
○ https://play.golang.org/p/glWZmm0gY6
● POINTER RECEIVER & POINTER VALUE
○ https://play.golang.org/p/pWFxsg6MSe
● POINTER RECEIVER & NON-POINTER VALUE
○ https://play.golang.org/p/G3lEy-4Mc8 ( code does not run )
○ this codes does run - notice the difference - method set determines
the INTERFACES that the type implements
■ https://play.golang.org/p/KK8gjsAWBZ
*/
