package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter integer: ")
	var input int
	fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(i int) {
	if i % 3 == 0 && i % 5 == 0 {
		fmt.Println("Fizzbuzz")
	} else if i % 3 == 0 {
		fmt.Println("Fizz")
	} else if i % 5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(i)
	}
}