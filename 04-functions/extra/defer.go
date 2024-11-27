package extra

import "fmt"

func deferFunction() {
	fmt.Println("-------- Functions - Defer --------")
	defer RunOne()
	RunTwo()
	RunThree()

	fmt.Println(calcAverage(8, 7))
}

func RunOne() {
	fmt.Println("Function 01")
}

func RunTwo() {
	fmt.Println("Function 02")
}

func RunThree() {
	fmt.Println("Function 03")
}

func calcAverage(n1, n2 int) bool {
	defer fmt.Println("Your overage will be calculated and returned")
	fmt.Println("Insert your grades to calcute the overage")
	if n1+n2/2 >= 6 {
		return true
	} else {
		return false
	}
}
