package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "peacock is an animal"
	fmt.Println("text")
	replacedtext := strings.Replace(text, "animal", "bird", 1)
	fmt.Println(replacedtext)

}
