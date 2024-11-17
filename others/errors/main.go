package main

import (
	"errors"
	"fmt"
)

func DoSomething() error {
	return errors.New("Returns an error as string")
}

func main() {
	fmt.Printf("%T", DoSomething()) // Prints *errors.errorString
	fmt.Println()
	fmt.Printf("%T", DoSomething().Error()) // Prints string
}
