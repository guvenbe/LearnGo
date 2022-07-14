//Go everything is pass by value
package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Flemming")
	fmt.Printf(" %s   %t", x, y)
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	//unfurling or spreading, xi... has to be the final parameter
	sum := sum("James", xi...)

	z := sum2("james")
	fmt.Println("The total is", z)

	fmt.Println("sum2=", sum)
	sayHello()
	sayHello("Sammy")
	sayHello("Sammy", "Jessica", "Drew", "Jamie")

	//Defer
	defer foo2()
	bar2()
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("hello, ", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `,says "Hello"`)
	b := false
	return a, b
}

func sum(y string, x ...int) int {
	fmt.Println(x) // this becomes a slice
	fmt.Println("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are adding", v, "to the total sum", sum)
	}
	fmt.Println("Sum=", sum)
	//fmt.Println("name=", name)
	return sum
}

func sayHello(names ...string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}
func sum2(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}
func foo2() {
	fmt.Println("Hello from foo2")
}

func bar2() {
	fmt.Println("bar2")
}
