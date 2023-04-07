package main

import (
	"fmt"
	"strings"
)

func main() {
	var Message string = "This Is My Book"
	SplittedStrings := strings.Split(Message, " ")
	fmt.Println(SplittedStrings)

}
