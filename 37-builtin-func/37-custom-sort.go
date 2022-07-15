package main

import (
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person
type ByName []Person

func (a ByAge) Len() int            { return len(a) }
func (a ByAge) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool  { return a[i].Age < a[j].Age }
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }
func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people)) // Byage(people) is  convertion
	fmt.Println(people)
	fmt.Println(people)
	sort.Sort(ByName(people)) // Byage(people) is  convertion
	fmt.Println(people)

	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Printf("%x", bs)

	loginPassword := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("you cant login")
		return
	}
	fmt.Println("success")
}
