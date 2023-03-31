package main

import "fmt"

func main() {
	var x int = 7
	var y int = 10
	if x == y {
		fmt.Printf("x is equal to y")
	} else if x > y {
		fmt.Printf("x is greater than y")
	} else {
		fmt.Printf("x is less than y")
	}
}
