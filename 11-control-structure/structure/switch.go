package structure

import "fmt"

func switchControl() {
	fmt.Println("-------- Control Structure -  SWITCH --------")

	fmt.Println("Day of the week: ", getDayWeek(10))

	fmt.Println("Day of the week: ", returnDayWeek(3))

	fmt.Println("Day of the week: ", dayOfTheWeek(5))

	fmt.Println("Testing fallThrougth: ", usingFallThrough(1))
}

func getDayWeek(day int) string {
	switch day {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Day not found"
	}
}

func returnDayWeek(day int) string {
	switch {
	case day == 1:
		return "Sunday"
	case day == 2:
		return "Monday"
	case day == 3:
		return "Tuesday"
	case day == 4:
		return "Wednesday"
	case day == 5:
		return "Thursday"
	case day == 6:
		return "Friday"
	case day == 7:
		return "Saturday"
	default:
		return "Day not found"
	}
}

func dayOfTheWeek(day int) string {
	var dayOfTheWeek string

	switch {
	case day == 1:
		dayOfTheWeek = "Sunday"
	case day == 2:
		dayOfTheWeek = "Monday"
	case day == 3:
		dayOfTheWeek = "Tuesday"
	case day == 4:
		dayOfTheWeek = "Wednesday"
	case day == 5:
		dayOfTheWeek = "Thursday"
	case day == 6:
		dayOfTheWeek = "Friday"
	case day == 7:
		dayOfTheWeek = "Saturday"
	default:
		dayOfTheWeek = "Day not found"
	}

	return dayOfTheWeek
}

func usingFallThrough(v int) string {
	var str string
	switch v {
	case 1:
		str = "I hope to return this string"
		fallthrough
	case 2:
		str = "Not today case 1! This is the string that will be returned"
	default:
		str = "The option doesn't exist"
	}

	return str
}
