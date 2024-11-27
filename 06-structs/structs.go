package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number int8
}

func main() {
	var user01 user
	fmt.Println(user01)
	fmt.Printf("The name of this user 01 is %s and your age %d \n", user01.name, user01.age)

	user01.name = "William"
	user01.age = 37
	fmt.Println(user01)
	fmt.Printf("The name of this user 01 is %s and your age %d \n", user01.name, user01.age)

	user02 := user{}
	fmt.Println(user02)
	fmt.Printf("The name of this user 02 is %s and your age %d \n", user02.name, user02.age)

	user03 := user{name: "Nadir", age: 47}
	fmt.Println(user03)
	fmt.Printf("The name of this user 03 is %s and your age %d \n", user03.name, user03.age)

	user04 := user{age: 47}
	fmt.Println(user04)
	fmt.Printf("This user04 only show us your age wich is %d \n", user04.age)

	user05 := user{name: "Jonathan"}
	fmt.Println(user05)
	fmt.Printf("This user05 only show us your name wich is %s \n", user05.name)

	address01 := address{"Tv Paulo Magalhães", 6}
	user06 := user{"Ramon Willim", 37, address01}
	fmt.Println("User 06 associated with an address", user06)

	user07 := user{"Roberto", 57, address{"In the somewhere in the world", 10}}
	fmt.Println("User 07 associated with an address", user07)

	var user08 user
	user08.name = "Wolverine"
	user08.age = 200
	user08.address.street = "Live in the X-MEN world"
	user08.address.number = 101
	fmt.Println("User 08 associated with an address", user08)
}
