package main

import (
	"fmt"
)

func main() {
	fmt.Println(plus(5, 10))
	hello()
	q, r := div(10, 3)
	fmt.Printf("%d %d\n", q, r)
	x, y := doSomething()
	fmt.Printf("%d %d\n", x, y)

	f := func(x, y int) int { return x + y }
	fmt.Println(f(2, 3))
	fmt.Printf("%T\n", f)

	f2 := later()
	fmt.Println(f2("Golang"))
	fmt.Println(f2("is"))
	fmt.Println(f2("awesome!"))

	f3 := integers()
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())
}

func plus(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello, World!")
	return
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func doSomething() (x, y int) {
	y = 5
	return
}

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
