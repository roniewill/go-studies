package extra

import "fmt"

func recoverThisFunc() {
	if r := recover(); r != nil {
		fmt.Println("Recovered successfully")
	}
}

func PanicRecoverFunc() {
	result := calcGradeAverage(6, 6)

	fmt.Println(result)

	fmt.Println("After runed this code")
}

func calcGradeAverage(n1, n2 int) bool {
	defer recoverThisFunc()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("The overage is exactly 6")
}
