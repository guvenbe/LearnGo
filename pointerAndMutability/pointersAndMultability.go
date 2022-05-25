package main

import "fmt"

type Artist struct {
	Name, Genre string
	Songs       int
}

func newReleasePassByValue(a Artist) int { //passing artist by value
	a.Songs++
	return a.Songs
}

func newReleasePassByReference(a *Artist) int { //passing by reference
	a.Songs++
	return a.Songs
}

func main() {
	me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s release their %d th song\n", me.Name, newReleasePassByValue(me))
	fmt.Printf("%s has to total of %d songs\n", me.Name, me.Songs)
	fmt.Println("\n")

	fmt.Printf("%s release their %d th song\n", me.Name, newReleasePassByReference(&me))
	fmt.Printf("%s has to total of %d songs\n", me.Name, me.Songs)
}
