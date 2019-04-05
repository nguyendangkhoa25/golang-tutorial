package main

import "fmt"

func main() {
	a, b := 5.67, 8.97

	fmt.Printf("Type of a %T, b %T\n", a, b)

	sum := a + b
	diff := a - b
	fmt.Println("Sum: ", sum, " Diff: ", diff)

	no1, no2 := 56, 89
	fmt.Println("Sum: ", no1+no2, " Diff: ", no1-no2)
}
