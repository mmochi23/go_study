package main

import (
	"fmt"
	"math"
)

func main() {
	var b bool
	b = true
	fmt.Println(b)

	var (
		n1 int
		n2 int8
		n3 int16
		n4 int32
		n5 int64
	)
	n1 = 9223372036854775807
	n2 = 127
	n3 = 32767
	n4 = 2147483647
	n5 = 9223372036854775807
	fmt.Printf("%d %d %d %d %d\n", n1, n2, n3, n4, n5)

	var (
		un1 uint
		un2 uint8
		un3 uint16
		un4 uint32
		un5 uint64
	)
	un1 = 9223372036854775807*2 + 1
	un2 = 127*2 + 1
	un3 = 32767*2 + 1
	un4 = 2147483647*2 + 1
	un5 = 9223372036854775807*2 + 1
	fmt.Printf("%d %d %d %d %d\n", un1, un2, un3, un4, un5)

	ln1 := 0777
	fmt.Printf("%d %o %T\n", ln1, ln1, ln1)
	ln2 := 0xF0FF
	fmt.Printf("%d %x %T\n", ln2, ln2, ln2)

	n := 257
	fmt.Printf("%T\n", n)
	ui8 := byte(n)
	fmt.Printf("%d %T\n", ui8, ui8)
	i64 := int64(n)
	fmt.Printf("%d %T\n", i64, i64)

	fmt.Printf("uint32 max value = %d\n", math.MaxUint32)
	fmt.Printf("float32 max value = %v\n", math.MaxFloat32)
	fmt.Printf("float64 max value = %v\n", math.MaxFloat64)

	fmt.Printf("%v\n", 1.3)
	fmt.Printf("%v\n", 1.03)
	fmt.Printf("%v\n", 1.003)
	fmt.Printf("%v\n", 1.0003)
	fmt.Printf("%v\n", 1.00003)
	fmt.Printf("%v\n", 1.000003)
	fmt.Printf("%v\n", 1.0000003)
	fmt.Printf("%v\n", 1.00000003)
	fmt.Printf("%v\n", 1.000000003)
	fmt.Printf("%v\n", 1.0000000003)
	fmt.Printf("%v\n", 1.00000000003)
	fmt.Printf("%v\n", 1.000000000003)
	fmt.Printf("%v\n", 1.0000000000003)
	fmt.Printf("%v\n", 1.00000000000003)
	fmt.Printf("%v\n", 1.000000000000003)
	fmt.Printf("%v\n", 1.0000000000000003)
	fmt.Printf("%v\n", 1.00000000000000003)
	fmt.Printf("%v\n", float32(1.3))
	fmt.Printf("%v\n", float32(1.03))
	fmt.Printf("%v\n", float32(1.003))
	fmt.Printf("%v\n", float32(1.0003))
	fmt.Printf("%v\n", float32(1.00003))
	fmt.Printf("%v\n", float32(1.000003))
	fmt.Printf("%v\n", float32(1.0000003))
	fmt.Printf("%v\n", float32(1.00000003))
	fmt.Printf("%v\n", float32(1.000000003))
	fmt.Printf("%v\n", float32(1.0000000003))
	fmt.Printf("%v\n", float32(1.00000000003))
	fmt.Printf("%v\n", float32(1.000000000003))
	fmt.Printf("%v\n", float32(1.0000000000003))
	fmt.Printf("%v\n", float32(1.00000000000003))
	fmt.Printf("%v\n", float32(1.000000000000003))
	fmt.Printf("%v\n", float32(1.0000000000000003))
	fmt.Printf("%v\n", float32(1.00000000000000003))

	s := `Start
Go
String
Sample\n
`
	fmt.Printf("%v", s)
}
