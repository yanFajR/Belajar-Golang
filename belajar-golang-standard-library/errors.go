package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("Not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id == "ryan" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("ryan")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown")
		}
	}
}

