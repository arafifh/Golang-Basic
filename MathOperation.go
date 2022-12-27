package main

import "fmt"

func main() {
	a := 23
	b := 24
	c := a + b
	d := b - a
	e := a * b
	f := a / b
	g := a % 5
	a += 100
	b++
	status := true

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(!status)
}