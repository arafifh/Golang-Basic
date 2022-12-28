package main

import "fmt"

func main() {
	person := map[string]string{
		"name":  "Rafif",
		"major": "Informatics Engineering",
	}

	person["country"] = "Indonesia"

	var1 := "city"
	var2 := "Surabaya"
	person[var1] = var2

	fmt.Println(person["name"], person["major"], person["country"], person[var1])

	person["name"] = "Hikmatiar"
	fmt.Println(person["name"], person["major"], person["country"], person[var1])

	delete(person, "country")
	fmt.Println(person)
}
