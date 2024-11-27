package extra

import "fmt"

func anonymousFunctions() {
	fmt.Println("-------- Functions - Anonymous function --------")

	func() {
		fmt.Println("Anonymous function just runed!")
	}()

	func(text string) {
		fmt.Println(text)
	}("Anonymous function with params")

	result := func(text string, number int) string {
		return fmt.Sprintf("Params: %s - %d", text, number)
	}("This is the text", 10)

	fmt.Println(result)
}
