package main

import (
	"fmt"
	"runtime"
)

func main() {
	ifSample(1)
	ifSample(2)
	ifSample(3)

	if x, y := 1, 2; x < y {
		fmt.Printf("x(%d) is less than y(%d)\n", x, y)
	}

	for {
		fmt.Println("Hello, World!")
		break
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Hello, World!!")
	}
	i := 1
	for i < 10 {
		fmt.Println("Hello, World!!!")
		i++
	}
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
		i++
	}
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		fmt.Printf("%d %s\n", i, s)
	}
	for i, s := range "Apple" {
		fmt.Printf("%d %x\n", i, s)
	}

	n := 3
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("others")
	}
	switch {
	case n > 0:
		fmt.Println("greater than 0")
	default:
		fmt.Println("others")
	}

	s := "A"
	switch s {
	case "A", "B":
		fmt.Println("A or B")
		fallthrough
	case "C", "D":
		fmt.Println("C or D")
		fallthrough
	default:
		fmt.Println("others")
	}

	var x interface{} = 3
	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("others")
	}

	fmt.Println("A")
	goto L
	fmt.Println("B")
L:
	fmt.Println("C")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 5 {
				goto DONE
			}
			fmt.Printf("i=%d j=%d\n", i, j)
		}
	}
DONE:

	runDefer()

	go sub()
	for i := 0; i < 100; i++ {
		fmt.Println("main loop")
	}
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())

	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("runtime error:%v\n", x)
		} else {
			fmt.Println("x is nil")
		}
	}()
	panic("panic")
	fmt.Println("end")
	defer fmt.Println("defer7")
}

func ifSample(x int) {
	if x == 1 {
		fmt.Println("xは1です")
	} else if x == 2 {
		fmt.Println("xは2です")
	} else {
		fmt.Println("xは1でも2でもないです")
	}
}

func runDefer() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	fmt.Println("done")
	defer func() {
		fmt.Println("defer4")
		fmt.Println("defer5")
		fmt.Println("defer6")
	}()
}

func sub() {
	for i := 0; i < 100; i++ {
		fmt.Println("sub loop")
	}
}

func init() {
	fmt.Println("first line")
}
