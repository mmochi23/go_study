package main

import (
	"fmt"
)

type MyError struct {
	Message string
	ErrCode int
}

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("%s [%s]", c.Model, c.Number)
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%d:%s", e.ErrCode, e.Message)
}

func RaiseError() error {
	return &MyError{Message: "私のエラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	s := []Stringify{
		&Person{Name: "Taro", Age: 21},
		&Car{Model: "PX512", Number: "XXX-0123"},
	}
	for _, s := range s {
		fmt.Println(s.ToString())
	}
}
