package main

import "fmt"

type Filter func(string)string

func sayWithFilter (name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}


func spamFilter(name string)string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayWithFilter("Ryan", spamFilter)

	filter := spamFilter

	sayWithFilter("anjing", filter)
}