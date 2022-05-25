package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

//type applies to both
func add2(x, y int) int {
	return x + y
}

func location(city string) (string, string) {
	var region, continent string
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

//named result
func location2(city string) (region, continent string) {
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"

	}
	return //returning region and continent
}
func main() {
	fmt.Println(add(42, 23))
	fmt.Println(add(60, 24))
	region, continent := location("NYC")
	fmt.Printf("Matt lives in %s, %s\n", region, continent)

	region2, continent2 := location2("Santa Monica")
	fmt.Printf("Matt live in %s, %s", region2, continent2)
}
