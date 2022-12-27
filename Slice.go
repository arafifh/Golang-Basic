package main

import "fmt"

func main() {
	var month = [...]string {
		"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des",
	}
	var slice = month[6:10]
	
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(append(slice, "blank"))

	slice[2] = "changed"
	fmt.Println(month)
}