package main

import (
	"fmt"
	"log"
	"os"
)

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

func file() {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("opened file", file.Name())
}
