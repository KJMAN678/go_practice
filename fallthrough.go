package main

import "fmt"

func main() {
	x := 4

	t := 0
	switch x {
	case 5:
		t += 5
		fallthrough
	case 4:
		t += 4
		fallthrough
	case 3:
		t += 3
		fallthrough
	case 2:
		t += 2
		fallthrough
	case 1:
		t += 1
	default:
		fmt.Println("範囲外です。")
		return
	}
	fmt.Println(t, "です。")
}
