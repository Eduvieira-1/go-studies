package main

import "fmt"

func main() {
	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	fmt.Println(len(myArray) - 1)

	for i, v := range myArray {
		fmt.Printf("O valor do indice %d é %d", i, v)
	}

}
