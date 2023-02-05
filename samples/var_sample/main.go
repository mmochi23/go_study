package main

import (
	"fmt"
)

var p = 100

func main() {
	var n int
	n = 5
	fmt.Println(n)

	var x, y, z int
	x, y = 1, 2
	fmt.Println(x, y, z)

	var (
		j, k int
		name string
	)
	j, k, name = 3, 4, "山田"
	fmt.Println(j, k, name)

	i := 1
	fmt.Printf("%T %d\n", i, i)
	b := true
	fmt.Printf("%T %v\n", b, b)
	f := 3.14
	fmt.Printf("%T %v\n", f, f)
	s := "abc"
	fmt.Printf("%T %v\n", s, s)
	one := one()
	fmt.Printf("%T %v\n", one, one)

	p = p + 1
	fmt.Println(p)
}

func one() float32 {
	return 1
}
