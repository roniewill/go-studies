package main

import "fmt"

func closure() func() {
	txt := "Text of closure function"

	_func := func() {
		fmt.Println(txt)
	}

	return _func
}

func main() {
	txt := "Into main function"

	fmt.Println(txt)

	_newFunc := closure()
	_newFunc()
}
