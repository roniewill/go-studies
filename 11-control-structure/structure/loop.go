package structure

import (
	"fmt"
)

func usingLoop() {
	fmt.Print("-------- Control Structure -  LOOP --------\n")

	// i := 0
	// fmt.Printf("Initial value: %d\n", i)
	// for i <= 5 {
	// 	time.Sleep(time.Second)
	// 	fmt.Printf("Current value is: %d\n", i)
	// 	i++
	// }

	// for j := 0; j <= 5; j++ {
	// 	time.Sleep(time.Second)
	// 	fmt.Printf("Initial value: %d\n", j)
	// }

	// for j := 0; j <= 10; j += 2 {
	// 	time.Sleep(time.Second)
	// 	fmt.Printf("Initial value: %d\n", j)
	// }

	// names := [3]string{"Mateus", "Marcos", "Lucas"}

	// for i, name := range names {
	// 	time.Sleep(time.Second)
	// 	fmt.Printf("Index: %d - Name: %s \n", i, name)
	// }

	// for i, char := range "SALVADOR" {
	// 	// it will show us the SAC table code related to the each letter of word
	// 	fmt.Println(i, char)
	// }

	// for i, char := range "SALVADOR" {
	// 	// this time it will show us each letter of word
	// 	fmt.Println(i, string(char))
	// }

	name := map[string]string{
		"name":   "William",
		"email":  "william@golang.com",
		"gender": "male",
	}

	for i, value := range name {
		fmt.Printf("%s: %s \n", i, value)
	}
}
