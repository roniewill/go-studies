package structure

import "fmt"

func ifInitControl() {
	fmt.Println("-------- Control Structure - IF ELSE and INIT --------")

	number01 := 15
	// using common 'IF' conditional
	if number01 > 15 {
		fmt.Println("Bigger than 15")
	} else {
		fmt.Println("Small or equal 15")
	}

	// using 'IF INIT' conditional
	// BTW: when we criate a variable using 'if init' that variable is limited to the scope
	if number02 := number01; number02 > 0 {
		fmt.Println("Number 02 is bigger than 0")
	}

}
