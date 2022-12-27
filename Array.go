package main

import "fmt"

func main() {
	var name [2]string
	name[0] = "Rafif"
	name[1] = "Hikmatiar"

	var age = [2]int8 {
		19, 20,
	}
	fmt.Println("My name is", name[0], name[1], "and i'm", age[0], "years old. I'm", age[1], "next year")

	fmt.Println(len(name))
}