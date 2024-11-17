package main

import "fmt"

func main() {
	variadicFunction(7, 1, 2, 3, 4, 5, 6, 7, 8)
}

func variadicFunction(num int, nums ...int) {
	fmt.Printf("Type of nums: %T\n", nums)
	for i, v := range nums {
		if v == num {
			fmt.Println("Found the num:", v, " at index: ", i)
		}
	}
}
