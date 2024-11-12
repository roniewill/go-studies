package main

import (
	"errors"
	"fmt"
)

func main() {
	// number01 := 23874
	// var number02 uint = 32800
	// var num8 int8 = 127
	// var num16 int16 = 32767
	// var num32 int32 = 2147483647
	// var num64 int64 = 9223372036854775807

	// fmt.Println("The maximum value for the int8 data type is: ", num8)
	// fmt.Println("The maximum value for the int16 data type is: ", num16)
	// fmt.Println("The maximum value for the int32 data type is: ", num32)
	// fmt.Println("The maximum value for the int64 data type is: ", num64)
	// fmt.Println("The value of the int data type is: ", number01)
	// fmt.Println("The value of the uint data type is: ", number02)
	// fmt.Printf("The type of number02 is: %T\n", number02)

	// var b byte = 'A'
	// fmt.Println(b)
	// fmt.Printf("%c\n", b)

	// var floatnum01 float32 = 123.987
	// var floatnum02 float64 = 12398665645.9870909
	// floatnum03 := 234234234.98762876234
	// fmt.Println(floatnum01)
	// fmt.Println(floatnum02)
	// fmt.Println(floatnum03)

	// var string01 string = "I'm a string type"
	// var string02 string = "I'm also a string type"
	// fmt.Println(string01, string02)

	// char := 'A'
	// fmt.Println("ASC table position: ", char)
	// fmt.Printf("Print the character: %c\n", char)

	// var bool01 bool
	// bool02 := true
	// fmt.Println(bool01)
	// fmt.Println(bool02)

	var erro error = errors.New("Internal error")
	fmt.Println(erro)
}
