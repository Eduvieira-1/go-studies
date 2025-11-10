package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(50, 10, 1230, 102394, 123445))

}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}
