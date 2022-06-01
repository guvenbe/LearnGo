package main

import (
	"fmt"
	"time"
)

var a int = 1

type hotdog int

var b hotdog = 3

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{}) //asserting y as map of interfaces
	if ok {
		z["updated_at"] = time.Now() //z now has the type map[string]interface
		z["car"] = "Chevy"
	}
}
func main() {
	foo := map[string]interface{}{
		"Matt": 42,
	}
	timeMap(foo)
	fmt.Println(foo)

	//conversion (not) casting
	a = int(b)
	b = hotdog(a)

}
