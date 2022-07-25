//Package word provides custom functions for working with words and strings
package word

import "strings"

//Use count returns the number times each word apperas in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count return the number words in a string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
