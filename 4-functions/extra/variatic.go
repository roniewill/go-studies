package extra

import "fmt"

func variatic() {
	fmt.Println("-------- Functions - Variatic --------")
	fmt.Println("Total: ", sumNumbers(1, 2, 33, 23, 12))
	writter("Something -", 10, 1, 2, 3, 45, 100)
}

func writter(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
		fmt.Println("Current value: ", total)
	}

	return total
}
