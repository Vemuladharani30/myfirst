package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "car engine"
	substring1 := "car"
	substring2 := "machine"
	result := strings.Contains(text, substring1)
	fmt.Println(result)
	result1 := strings.Contains(text, substring2)
	fmt.Println(result1)
}
