package main

import "testing"

type test struct {
	data   []int
	answer int
}

func TestMySum(t *testing.T) {
	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}
	for _, v := range tests {
		x := MySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func TestMySub(t *testing.T) {

	x := MySubraction(3, 2)
	if x != -5 {
		t.Error("Expected", -5, "Got", x)
	}
}
