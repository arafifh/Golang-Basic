package main

import "fmt"

func main() {
	var university string
	var years = 2
	major := "Informatics Engineering"
	var (
		name string = "Rafif Hikmatiar"
		age int8 = 19
	)
	university = "Sepuluh Nopember Institute of Technology"

	fmt.Println("My name is", name, "i'm", age, "years old. Now i study in", university, "and this is my", years, "nd years, my major is", major)
}