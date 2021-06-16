package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x := scanner.Text()

	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Error!")
		return
	}
	p := float64(n)
	fmt.Println(int(p * 1.1))
}
