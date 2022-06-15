package main

import (
	"fmt"
)

func main() {
	//x:=type{values} //composite literal
	x := []int{4, 5, 6, 7, 8}
	//slice allows you to group together VALUES of the same TYPE
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Printf("index:%d, value:%v\n", i, v)
	}

	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i := 0; i < len(x); i++ {
		fmt.Printf("%d\t%d\n", i, x[i])
	}

	x = append(x, 77, 88, 99, 1024)
	fmt.Println(x)

	y := []int{234, 344, 555, 999}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[0:2], x[3:]...)
	fmt.Println(x)

	z := make([]int, 10, 12)
	fmt.Println(y)
	fmt.Println("--------------------")
	fmt.Println(len(z))
	fmt.Println(cap(z))
	z[0] = 1
	z[9] = 10
	//z[10] = 12))
	fmt.Println("--------------------")
	fmt.Println(len(z))
	fmt.Println(cap(z))
	z = append(z, 13)
	z = append(z, 13)
	z = append(z, 13)

	fmt.Println("--------------------")
	fmt.Println(len(z))
	fmt.Println(cap(z))

	jb := []string{"james", "bond", "choclate", "martini"}
	fmt.Println(jb)
	mp := []string{"miss", "Moeypanny", "strawberry", "Hazlenut"}
	fmt.Println(jb)
	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
