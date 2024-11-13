package main

import "fmt"

func main() {
	result := sum(5, 3)
	fmt.Println(result)

	var f = func(txt string) string {
		return txt
	}

	txt01 := f("some text")
	fmt.Println(txt01)

	result01, result02 := mathCalc(20, 12)
	fmt.Println(result01, result02)

	result03, _ := mathCalc(20, 5)
	fmt.Println(result03)

	_, result04 := mathCalc(20, 5)
	fmt.Println(result04)
}

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalc(n1, n2 int8) (int8, int8) {
	result_sum := n1 + n2
	result_sub := n1 - n2
	return result_sum, result_sub
}
