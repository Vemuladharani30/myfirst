package main

import (
	"fmt"
)

type phone struct {
	ram    int
	colour string
	price  int
}

func main() {
	x := phone{ram: 16, colour: " black ", price: 6500}
	fmt.Println("%v", x)
}
