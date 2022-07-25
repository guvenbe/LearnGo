package main

import (
	"LearnGo/42-benchmark/excercise2/quote"
	"LearnGo/42-benchmark/excercise2/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
