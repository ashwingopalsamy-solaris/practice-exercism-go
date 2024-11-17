package main

import "fmt"

func main() {
	fmt.Print("Hello Print", 21)
	fmt.Println("Hello Println", 23)
	fmt.Printf("Hello %d ",23)
	fmt.Println(fmt.Sprintln("Hello %d ",23))
	fmt.Println("Hi")
}
