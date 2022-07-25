package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	dogYears := Years(10)
	if dogYears != 70 {
		t.Error("got", dogYears, "expected", 70)
	}
}

func TestYearsTwo(t *testing.T) {

	dogYears := YearsTwo(10)
	if dogYears != 70 {
		t.Error("got", dogYears, "expected", 70)
	}
}
func ExampleYears() {

	fmt.Println(Years(10))
	// Output:
	// 70
}
func ExampleYearsTwo() {

	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
