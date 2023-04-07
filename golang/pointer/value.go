package main

import (
	"fmt"
)

func main() {
	var ptr int
	ptr = 90
	va := &ptr
	fmt.Println(*va)

}
