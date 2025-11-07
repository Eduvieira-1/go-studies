package main

import "fmt"

func main() {

	salaries := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}

	fmt.Println(salaries["Wesley"])

	// salary_1 := make(map[string]int)
	// salary_2 := map[string]int{}
	// salary_1["Wesley"] = 1000

	for name, salary := range salaries {
		fmt.Printf("O salario de %s é %d\n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("O salario é %d\n", salary)
	}
}
