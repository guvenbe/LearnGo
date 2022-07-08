package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"Miss MoneyPenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok)
	}

	m["todd"] = 34
	for k, v := range m {
		fmt.Printf("%s\t%d\n", k, v)
	}

	//slice
	xi := []int{1, 2, 3, 4, 5, 6}
	for i, v := range xi {
		fmt.Printf("%d\t%d\n", i, v)
	}

	k := "james"
	if v, ok := m[k]; ok {
		fmt.Println(v)
		delete(m, k)
	} else {
		fmt.Printf("%s,  not found", k)
	}
	fmt.Println(m)
}
