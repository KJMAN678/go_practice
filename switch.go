package main

import (
	"fmt"
)

func main() {
	x := 1
	switch x {
	case 1:
		fmt.Println("1でした")
	case 2:
		fmt.Println("2かもしれません")
	default:
		fmt.Println("わかりませんでした")
	}
}
