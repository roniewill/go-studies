package main

import "fmt"

func main() {
	// ## assigning values ​​by copy
	// var pointer01 int = 100
	// pointer02 := pointer01
	// fmt.Println(pointer01, pointer02)

	// pointer01++
	// fmt.Println(pointer01, pointer02)

	// ## assigning values ​​by memory reference AKA pointers
	var number int = 30
	var pointer *int
	fmt.Println("Number is: ", number)
	fmt.Println("Memory location is: ", pointer)

	pointer = &number
	fmt.Println("Number is: ", number)
	fmt.Println("Memory location is: ", pointer)
	fmt.Println("Volue in the memory location is: ", *pointer)

	number++
	fmt.Println("Number is: ", number)
	fmt.Println("Memory location is: ", pointer)
	fmt.Println("Volue in the memory location is: ", *pointer)
}
