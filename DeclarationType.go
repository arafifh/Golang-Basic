package main

import "fmt"

func main() {
	type status bool
	var legal status = true

	fmt.Println(legal)
}