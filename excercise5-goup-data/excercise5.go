package main

import "fmt"

//Array
func main() {
	var x [5]int

	for i := 0; i < 5; i++ {
		x[i] = i
	}

	for k, v := range x {
		fmt.Printf("%d - %d\n", k, v)

	}

	y := [5]int{44, 55, 22, 37, 90}
	for i, v := range y {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", y)

	//slice

	z := []int{12, 22, 66, 90, 4, 9, 5, 89, 33, 7}
	for i, v := range z {
		fmt.Printf("%d - %d\n", i, v)
	}
	fmt.Printf("%T\n", z)

	w := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(w[:5])
	fmt.Println(w[5:])
	fmt.Println(w[2:7])
	fmt.Println(w[1:6])
	w = append(w, 52)
	fmt.Println(w)
	w = append(w, 53, 54, 55)
	fmt.Println(w)
	a := []int{56, 57, 58, 59, 60}
	w = append(w, a...)
	fmt.Println(w)

	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	b = append(b[:3], b[6:]...)
	fmt.Println(b)

	t := make([]string, 50, 50)
	fmt.Println("first time")
	fmt.Println(len(t))
	fmt.Println(cap(t))
	t = append(t, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println("second time")
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	for i := 0; i < len(t); i++ {
		fmt.Println(i, t[i])
	}

	// CORRECT CODE IS BELOW
	s := make([]string, 50, 50)
	fmt.Println("third time WITH Y")
	fmt.Println(len(s))
	fmt.Println(cap(s))

	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	for i, v := range states {
		s[i] = v
	}
	fmt.Println("fourth time WITH Y")
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
	m := [][]string{{"james", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloo James"}}

	fmt.Println(m)
	for i, row := range m {
		fmt.Printf("\nrow: %d\n", i)
		for col, value := range row {
			fmt.Printf("\t col%d\t %v\n", col, value)
		}
	}
	myMap := map[string][]string{
		"bond_james":      []string{`Shaken not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literatre`, "Computer Science"},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	//fmt.Println(myMap)
	for k, v := range myMap {

		fmt.Printf("Row: %s\n", k)
		for i, val := range v {
			fmt.Printf("\t index %d \t %v\n", i, val)
		}
	}
	myMap["Q"] = []string{"Don't touch that double O 7", "Inventing", "science"}

	fmt.Println("Added Q\n\n")
	for k, v := range myMap {

		fmt.Printf("Row: %s\n", k)
		for i, val := range v {
			fmt.Printf("\t index %d \t %v\n", i, val)
		}
	}
	if _, ok := myMap["Q"]; ok {
		delete(myMap, "Q")
	}

	fmt.Printf("...Delete Q\n\n")
	for k, v := range myMap {

		fmt.Printf("Row: %s\n", k)
		for i, val := range v {
			fmt.Printf("\t index %d \t %v\n", i, val)
		}
	}
}
