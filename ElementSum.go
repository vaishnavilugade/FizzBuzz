package main

import "fmt"

func main() {
	var size int
	var elements int
	var sum int = 0
	fmt.Print("Enter size:")
	fmt.Scan(&size)
	fmt.Print("Enter elements:")
	for i := 1; i <= size; i++ {
		fmt.Scan(&elements)
		sum = sum + elements
	}

	fmt.Print("sum of elemets:", sum)

}
