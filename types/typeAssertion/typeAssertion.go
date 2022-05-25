package main

import (
	"fmt"
	"time"
)

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
}
