package main

import (
	"fmt"
	"strconv"
)

type intp int

func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (num intp) PrimeFactor() []int {
	ar := []int{}
	x := int(num)
	n := 2
	for x > n {
		if x%n == 0 {
			x /= n
			ar = append(ar, n)
		} else {
			if n == 2 {
				n++
			} else {
				n += 2
			}
		}
	}
	ar = append(ar, x)
	return ar
}

func main() {
	var message string
	fmt.Printf("type number:")
	fmt.Scanf("%s", &message)
	n, _ := strconv.Atoi(message)
	x := intp(n)
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x *= 2
	x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
}
