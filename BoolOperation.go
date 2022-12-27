package main

import "fmt"

func main() {
	testResult := 90
	attendance := 50
	
	// true -> passed
	// false -> !passed
	fmt.Println(testResult >= 80)
	fmt.Println(attendance >= 70)
	fmt.Println(testResult >= 80 && attendance >= 70)
}