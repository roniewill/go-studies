package extra

import "fmt"

func invertSignal(number int) int {
	return number * -1
}

func invertSignalByPointer(number *int) {
	// *number = *number * -1
	*number = *number + 1
}

func TestPointerFunc() {
	number := 20
	fmt.Println("Inverted number by value copied")
	signalInverted := invertSignal(number)
	fmt.Println(signalInverted)

	fmt.Println("The value is the same yet")
	fmt.Println(number)

	fmt.Println("Now this value has been modified using pointer")
	invertSignalByPointer(&number)
	fmt.Println(number)
}
