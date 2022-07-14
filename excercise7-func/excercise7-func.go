package main

import (
	"fmt"
	"math"
)

type square struct{ width, length float64 }
type circle struct{ radius float64 }

func (s square) area() float64 {
	return s.width * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface{ area() float64 }

func info(s shape) {

	switch s.(type) {
	case square:
		fmt.Printf("SQUARE Area: %f\n ", s.(square).area())
	case circle:
		fmt.Printf("CIRCLE Area: %f\n", s.(circle).area())
	}

}

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Printf("My name is %s %s and I am %d years old", p.first, p.last, p.age)
}

func retFunc() func() int {
	return func() int {
		fmt.Println("returned func")
		return 42
	}
}

func even(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

func check(x int, f func(x int) bool) bool {
	return f(x)
}

func sliceOperation(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

func lastFirstAdder(xi []int) int {
	if len(xi) == 0 {
		return 0
	}
	if len(xi) == 1 {
		return 1
	}

	return xi[0] + xi[len(xi)-1]
}

func doubler() func() int {
	var x int = 1
	return func() int {
		x *= 2
		return x
	}
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
	x, s := bar()
	fmt.Println(x, s)
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo2(xi...))
	fmt.Println(bar2(xi))
	defer foo()
	foo3()
	bar()
	p := person{first: "James", last: "Bond", age: 32}
	p.speak()

	fmt.Println("----------------------\n")
	sq := square{
		width:  2,
		length: 2}
	fmt.Println("----------------------\n")
	fmt.Println("square Area =", sq.area())
	ci := circle{
		radius: 2,
	}
	fmt.Println("circle Area =", ci.area())
	fmt.Println("----------------------\n")
	info(sq)
	info(ci)
	fmt.Println("----------------------\n")
	fmt.Println("----------------------\n")
	f := func() {
		for i := 0; i < 4; i++ {
			fmt.Println(i)
		}
	}
	f()
	fmt.Printf("%T\n", f)
	g := f
	g()

	fmt.Printf("%T\n", g)

	fmt.Println("----------Return Func------------\n")
	fmt.Println(retFunc()())

	fmt.Println("----------callback------------\n")
	fmt.Println(check(2, even))
	fmt.Println(sliceOperation(lastFirstAdder, []int{1, 2, 3, 4, 5, 6}))
	d1 := doubler()
	d2 := doubler()
	fmt.Println(d1())
	fmt.Println(d1())
	fmt.Println(d1())
	fmt.Println(d1())
	fmt.Println(d2())
	fmt.Println(d2())

}

func foo() int {
	fmt.Println("foo ran")
	return 1

}

func bar() (int, string) {
	fmt.Println("bar ran")
	return 1, "hello"
}

func foo2(i ...int) int {
	var sum int
	for _, v := range i {
		sum += v
	}

	return sum
}
func bar2(i []int) int {
	var sum int
	for _, v := range i {
		sum += v
	}

	return sum
}

func foo3() {
	defer func() {
		fmt.Println("Hello, from DEFER")
	}()
	fmt.Println("This ran")
}
