package main

import (
	"fmt"
)

func main() {
	quality := [2]string{"good", "bad"}
	noun := [3]string{"place", "person", "name"}
	for i := 0; i < len(quality); i++ {
		for j := 0; j < len(noun); j++ {
			fmt.Println(quality[i], noun[j])
		}

	}

}
