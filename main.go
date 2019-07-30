package main

import "fmt"

func main() {
	fmt.Println("Welcome to the dashboard!")
}

func deferred() {
	for range []int{1, 2, 3} {
		defer func() {
			fmt.Printf("hi")
		}()
	}
}
