package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Point struct {
	X, Y int
}

type Feed struct {
	Name   string `json:"feed_name"`
	Amount uint   `json:"feed_amount"`
}

type Animal struct {
	Name string `json:"animal_name"`
	Feed Feed   `json:"feed"`
}

func NewAnimal(name string, feed_name string) *Animal {
	a := new(Animal)
	a.Name = name
	a.Feed.Name = feed_name
	return a
}

func (a *Animal) feed() {
	a.Feed.Amount = 15
}

func main() {
	var pt Point = Point{Y: 3, X: 8}
	fmt.Println(pt.X, pt.Y)
	pt.X = 10
	pt.Y = 5
	fmt.Println(pt.X, pt.Y)

	a := NewAnimal("Monkey", "Banana")
	fmt.Println(a)
	a.feed()
	fmt.Println(a)

	as := make([]*Animal, 5)
	as[0] = NewAnimal("Elephant", "Glass")
	as[1] = NewAnimal("Rabbit", "Carrot")
	for _, a := range as {
		if a != nil {
			a.Feed.Amount = 10
		}
		fmt.Println(a)
	}

	m1 := map[int]*Animal{
		1: {Name: "Monkey", Feed: Feed{Name: "Banana", Amount: 10}},
	}
	fmt.Println(m1[1])

	t := reflect.TypeOf(Feed{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}

	bs, err := json.Marshal(a)
	if err == nil {
		fmt.Println(string(bs))
	}
}
