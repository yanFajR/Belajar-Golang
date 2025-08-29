package main

import "fmt"

func main() {
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Ryan"
	newSlice[1] = "Fajri"

	fmt.Println(newSlice[0])
	fmt.Println(newSlice[1])
	
	newSlice2 := append(newSlice, "Alifah")

	fmt.Println(newSlice2)

	targetSlice := make([]string, len(newSlice), cap(newSlice))
	copy(targetSlice, newSlice)

	fmt.Println(targetSlice)
}