package main

import "fmt"

func main() {
	fmt.Println("Running the main function after init function")
}

func init() {
	fmt.Println("Running the init function before main function")
}
