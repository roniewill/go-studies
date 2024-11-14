package main

import (
	"fmt"
)

type person struct {
	name   string
	age    uint8
	gender string
}

type student struct {
	person
	course  string
	college string
}

func main() {
	var student01 student
	student01.name = "William"
	student01.age = 37
	student01.gender = "male"
	student01.course = "Systems Analysis"
	student01.college = "UniCesumar"

	fmt.Println("Student 01 info: ", student01)

	person02 := person{"Jhonn", 24, "male"}
	fmt.Println("Person's info: ", person02)

	student02 := student{person02, "Rider", "Life's school"}
	fmt.Println("Student 02 info: ", student02)

	student02.age = 40
	fmt.Println(student02.age)
}
