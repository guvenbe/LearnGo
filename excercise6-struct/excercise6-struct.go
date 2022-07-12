package main

import "fmt"

type person struct {
	first, last    string
	iceCreamFlavor []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	fmt.Println()
	p1 := person{
		first:          "James",
		last:           "Bond",
		iceCreamFlavor: []string{"vanila", "choclate"},
	}
	fmt.Println(p1)

	p2 := person{
		first:          "Max",
		last:           "Smart",
		iceCreamFlavor: []string{"mint", "strawberry"},
	}

	fmt.Printf("Frist %s, Last %s\n", p2.first, p2.last)
	for _, v := range p2.iceCreamFlavor {

		fmt.Printf("\t %v", v)
	}
	m := map[string]person{"jb": p1, "sm": p2}
	fmt.Println("\n\n", m)

	for k, v := range m {
		fmt.Printf("\n %s --> %s, %s", k, v.first, v.last)
		for i, val := range v.iceCreamFlavor {
			fmt.Printf("\n\t %d --> %v", i, val)
		}

	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	chevy := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}
	cadillac := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(chevy)
	fmt.Println(chevy.color)
	fmt.Println(cadillac)
	fmt.Println(cadillac.doors)

	usa := struct {
		state    map[string]string
		mountain []string
	}{
		state: map[string]string{
			"MI": "Michigan",
			"CO": "Colorado",
		},
		mountain: []string{"Porkupine", "Rocky"},
	}

	fmt.Println(usa)
	for k, v := range usa.state {
		fmt.Printf("%s - %s\n", k, v)
	}
	for i, v := range usa.mountain {
		fmt.Printf("%d - %s\n", i, v)
	}
}
