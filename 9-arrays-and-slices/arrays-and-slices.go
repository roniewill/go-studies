package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ## Arrays
	fmt.Println("-------- Arrays --------")
	var array01 [5]string
	fmt.Println(array01)

	array02 := [3]int{}
	fmt.Println(array02)

	array03 := [4]int{10, 23, 40, 52}
	fmt.Println(array03)

	array04 := [...]int{10, 12, 14, 18, 16}
	// array04[5] = 17 // BTW: that assignment is not allowed once the array size already defined
	fmt.Println(array04)

	// ## Slices
	fmt.Println("-------- Slices --------")
	slice01 := []int{02, 8, 12}
	fmt.Println("Slice 01: ", slice01)
	slice01[0] = 03
	fmt.Println("Slice 01: ", slice01)

	slice01 = append(slice01, 13) // we can add one or more value in our slice
	fmt.Println("Slice 01: ", slice01)
	slice01 = append(slice01, 9, 17, 22) // we can add one or more value in our slice
	fmt.Println("Slice 01: ", slice01)

	slice02 := array04[1:4]
	fmt.Println("Slice 02: ", slice02)

	array04[3] = 250
	fmt.Println("Slice 02: ", slice02)

	// ## Checking types
	fmt.Println("Array type: ", reflect.TypeOf(array03))
	fmt.Println("Slice type: ", reflect.TypeOf(slice01))

	// ## Internal Arrays
	fmt.Println("-------- Internal Arrays - Arrays and Slices --------")
	slice03 := make([]float32, 3, 5)
	fmt.Println("Slice 03: ", slice03)

	slice03[2] = 150.40
	fmt.Println("Slice 03: ", slice03)

	fmt.Println("Length: ", len(slice03))
	fmt.Println("capacity: ", cap(slice03))

	slice03 = append(slice03, 13.3, 14.5, 18.33)
	fmt.Println("Slice 03 UPDATED: ", slice03)
	fmt.Println("Length: ", len(slice03))
	fmt.Println("capacity: ", cap(slice03))

	slice04 := make([]float32, 3)
	fmt.Println("Slice 04: ", slice04)
	fmt.Println("Length: ", len(slice04))
	fmt.Println("capacity: ", cap(slice04))

	slice04 = append(slice04, 4, 4.5, 6.8)
	fmt.Println("Slice 04: ", slice04)
	fmt.Println("Length: ", len(slice04))
	fmt.Println("capacity: ", cap(slice04))
}
