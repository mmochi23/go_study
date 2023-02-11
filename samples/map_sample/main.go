package main

import (
	"fmt"
)

func main() {
	m1 := make(map[int]string)
	m1[1] = "US"
	m1[81] = "Japan"
	m1[86] = "China"
	m1[81] = "日本"
	fmt.Println(m1)

	m2 := map[int]string{1: "Taro", 2: "Jiro", 3: "Saburo"}
	fmt.Println(m2)
	fmt.Println(m2[1])
	fmt.Println(m2[4])
	s1, ok1 := m2[1]
	s2, ok2 := m2[4]
	fmt.Printf("%s %v\n", s1, ok1)
	fmt.Printf("%s %v\n", s2, ok2)
	if _, ok := m2[1]; ok {
		fmt.Println("m2[1]が存在する")
	}

	for k, v := range m2 {
		fmt.Printf("key=%v value=%v\n", k, v)
	}

	delete(m2, 2)
	fmt.Println(m2)
}
