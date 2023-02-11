package main

import (
	"fmt"
)

const ONE = 1

const (
	X = 1
	Y
	Z  uint64 = X + Y
	S1        = "あ"
	S2
	S3 = S1 + "：" + S2
)

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}

func main() {
	x, y := one()
	fmt.Printf("%d %d\n", x, y)

	fmt.Printf("%d %d %d\n", X, Y, Z)
	fmt.Printf("%T %T %T\n", X, Y, Z)
	fmt.Printf("%s %s %s\n", S1, S2, S3)
}
