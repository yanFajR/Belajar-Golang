package main

import "fmt"

func main() {
	var buah [3]string

	buah[0] = "Pisang"
	buah[1] = "Anggur"
	buah[2] = "Jeruk"

	fmt.Println(buah[0], buah[1], buah[2])

	var nilai = [3]int {
		1, 2, 3,
	}

	fmt.Println(nilai[0], nilai[1], nilai[2])
}