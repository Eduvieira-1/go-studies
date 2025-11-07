package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)

	fmt.Println(sum(10, 40))
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}
