package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "dharani"
	string2 := "dharani vemula"
	string3 := "dharani"
	fmt.Println(strings.Compare(string1, string2))
	fmt.Println(strings.Compare(string2, string3))
	fmt.Println(strings.Compare(string3, string1))

}
