package main

import (
	"fmt"
)

func main() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	//fmt.Println("Value of c1: ", c1, " and c2: ", c2)

	cadd := c1 + c2

	fmt.Println("Sum: ", cadd)
	cmul := c1 * c2
	fmt.Println("Product: ", cmul)
}
