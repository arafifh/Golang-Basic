package main

import "fmt"

func main() {
	var value int32 = 2434254
	var valueConv int64 = int64(value)
	var valueConv2 int8 = int8(valueConv)

	fmt.Println(valueConv)
	fmt.Println(valueConv2)
}