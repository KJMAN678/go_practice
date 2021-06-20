package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x string
	fmt.Scanf("%s", &x)
	t := 0
	n, err := strconv.Atoi(x)
	if err != nil {
		goto err
	}
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!")
}
