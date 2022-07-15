package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type person struct {
	First, Last string
	Age         int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs2 := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs2)

	type person struct {
		First string `json:"First"`
		Last  string `json:"Last"`
		Age   int    `json:"Age"`
	}

	var people2 []person

	err = json.Unmarshal(bs2, &people2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people2)

	for i, v := range people2 {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	fmt.Fprintln(os.Stdout, "Hello world")
	io.WriteString(os.Stdout, "Hello word")

	xi := []int{4, 2, 5, 61, 12, 45, 13, 44, 3, 2, 1}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println("---------------------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
