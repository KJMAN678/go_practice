package main

import (
	"fmt"
)

func main() {
	x := 14

	t := 0
	c := 0
	for {
		c++
		if c%2 == 1 {
			continue
		}
		if c > x {
			break
		}
		t += c
	}
	fmt.Println(t, "です。")
}
