package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
	case (4 == 4):
		fmt.Println("also true, it does not prints")
	default:
		fmt.Println("this is default")
	}

	fmt.Println("========================")
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, it does not prints")
		fallthrough
	case (67 == 89):
		fmt.Println("alsoc print with fallthrough")
	}

	n := "Bond"
	switch n {
	case "MoneyPenny":
		fmt.Println("MoneyPenny")
	case ("Bond"):
		fmt.Println("Bond")
	case ("Q"):
		fmt.Println("Q")
	case "Gold Finger":
		fmt.Println("Gold Finger")
	default:
		fmt.Println("this is default")
	}

	x := "Bond"

	switch x {
	case "MoneyPenny", "bond", "DrNo":
		fmt.Println("MoneyPenny or bond or DrNo")
	case ("Bond"):
		fmt.Println("Bond")
	case ("Q"):
		fmt.Println("Q")
	case "Gold Finger":
		fmt.Println("Gold Finger")
	default:
		fmt.Println("this is default")
	}
}
