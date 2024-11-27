package extra

import "fmt"

func recursiveFunctions() {
	fmt.Println("-------- Functions - Recursive functions --------")

	positions := uint(20)

	for i := uint(0); i <= positions; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
