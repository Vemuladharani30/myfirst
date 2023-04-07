package main

import (
	"fmt"
)

func myfunction(x int, y int) (result int) {
	result = x + y
	return
}
func main() {
	fmt.Println(myfunction(4, 3))
}
