package extra

import "fmt"

func returnNamed() {
	fmt.Println("-------- Functions - Return Named --------")
	sum, subt := MathCalculation(20, 18)
	fmt.Println(sum, subt)
}

func MathCalculation(n1, n2 int) (sum int, subt int) {
	sum = n1 + n2
	subt = n1 - n2
	return
}
