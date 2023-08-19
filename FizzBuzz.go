
package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz\n")
		} else if i%5 == 0 {
			fmt.Print("Buzz\n")
		} else if i%3 == 0 {
			fmt.Print("Fizz\n")
		} else {
			fmt.Print(i, "\n")
		}
	}
}
