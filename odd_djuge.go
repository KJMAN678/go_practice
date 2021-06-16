package main

import "fmt"

func main() {
	x := 10
	switch {
	case x%2 == 0:
		fmt.Println("偶数です。")
	case x%2 == 1:
		fmt.Println("期数です。")
	}

}
