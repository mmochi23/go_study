package main

import (
	"fmt"
)

func main() {
	var i int
	p := &i
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %v %T\n", p, *p, p)
	i = 5
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %v %T\n", p, *p, p)
	*p = 10
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %v %T\n", p, *p, p)

	inc(p)
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %v %T\n", p, *p, p)
	inc(&i)
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %v %T\n", p, *p, p)

}

func inc(p *int) {
	*p++
}
