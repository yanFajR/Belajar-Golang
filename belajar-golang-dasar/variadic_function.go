package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {

	nomor := []int{10,10,10,10}
	fmt.Println(sum(1,2,3,4,5))
	fmt.Println(sum(nomor...))

	tes := sum

	fmt.Println(tes(1,2,3,4,5,6,7,8))

}