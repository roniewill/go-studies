package main

import "fmt"

func main() {
	fmt.Println("-------- Maps --------")

	// Maps type of string
	user := map[string]string{
		"name":     "Ramon William",
		"lastname": "Gomes",
	}
	fmt.Println("User: ", user)
	fmt.Println("User - name: ", user["name"])
	fmt.Println("User - lastname: ", user["lastname"])

	// Maps that has as a key string and value a map
	user02 := map[string]map[string]string{
		"fullname": {
			"name":     "Ramon William",
			"lastname": "Gomes",
		},
		"person_info": {
			"profession": "Software Engineer",
			"degree":     "Computer Sciencie",
		},
		"todelete": {
			"one": "one",
			"two": "two",
		},
		"todelete02": {
			"one": "one",
			"two": "two",
		},
	}

	fmt.Println("User 02: ", user02)

	fmt.Println("---------------------------------")

	delete(user02, "todelete")

	fmt.Println("User 02: ", user02)

	fmt.Println("---------------------------------")

	delete(user02, "todelete02")

	fmt.Println("User 02: ", user02)

	fmt.Println("---------------------------------")

	user02["devlangs"] = map[string]string{
		"first":  "Golang",
		"second": "Ruby",
		"third":  "Javascript",
	}

	fmt.Println("User 02: ", user02)

	fmt.Println("---------------------------------")

	user03 := map[string]map[int8]string{
		"first": {
			1: "one",
			2: "two",
			3: "three",
		},
		"second": {
			1: "four",
			2: "five",
		},
	}

	fmt.Println("User 03: ", user03)
	fmt.Println("User 03: ", user03["first"])
	fmt.Println("User 03: ", user03["second"])
	fmt.Println("User 03: ", user03["first"][2])
	fmt.Println("User 03: ", user03["second"][2])
}
