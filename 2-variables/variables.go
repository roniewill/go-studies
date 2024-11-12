package main

import "fmt"

func main() {
	// var var01 string = "Variable one"
	// var02 := "Variable two"
	// var03 := 3000
	// fmt.Println(var01)
	// fmt.Println(var02)
	// fmt.Println(var03)

	var (
		var04 string = "Variable four"
		var05 string = "Variable five"
	)

	fmt.Println(var04, var05)

	var06, var07 := "Variable six", "50.500"

	var06, var07 = var04, var05

	fmt.Println(var06, var07)

	// const const01 string = "Constant 01"
	// fmt.Println(const01)

	// const const02 = "Constant 02"
	// fmt.Println(const02)

	// const const03 = 2259710620
	// fmt.Println("Value of Constant 03: ", const03)

}
