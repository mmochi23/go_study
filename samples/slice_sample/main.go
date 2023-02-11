package main

import (
	"fmt"
)

func main() {
	s := make([]int, 5, 10)
	s[3] = 3
	fmt.Printf("len=%d cap=%d value=%v\n", len(s), cap(s), s)
	s = append(s, 5, 6, 7)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s), cap(s), s)
	s = append(s, 8, 9, 10)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s), cap(s), s)

	a := [5]int{1, 2, 3, 4, 5}
	b := a[0:2]
	fmt.Printf("%T %v\n", b, b)

	s = append(s, b...)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s), cap(s), s)

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{11, 12}
	s3 := copy(s1, s2)
	fmt.Printf("%v %v %v\n", s1, s2, s3)

	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := a1[2:4:6]
	fmt.Printf("%v len=%d cap=%d\n", a2, len(a2), cap(a2))

	for i, v := range a1 {
		fmt.Printf("%d %d\n", i, v)
	}

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(s1...))

	p1 := [3]int{1, 2, 3}
	pow(p1)
	fmt.Println(p1)
	p2 := []int{1, 2, 3}
	pows(p2)
	fmt.Println(p2)
}

func sum(s ...int) int {
	fmt.Printf("%T\n", s)
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func pow(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func pows(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}
