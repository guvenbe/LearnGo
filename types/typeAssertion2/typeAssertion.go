package main

import "fmt"

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	str, ok := value.(string)
	if ok { //this is the "ok" idiom
		fmt.Println(str)
	}
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}
func main() {
	s := &fakeString{"Ceci n'est pas un string"}
	printString(s)
	printString("Hello Gophers")
}
