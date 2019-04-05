package main

import "fmt"

func main() {
	var myAge int //Variable declaration
	fmt.Println("Hello world!")
	fmt.Println("my age is: ", myAge)
	myAge = 15
	fmt.Println("my new age is: ", myAge)
	myAge = 30
	fmt.Println("My new age is: ", myAge)

	name, age := "Steven", 29
	fmt.Println("my name is:", name, " age is", age)

}
