package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("input data: ")
	scanner.Scan()
	x := scanner.Text()

	ar := strings.Split(x, " ")
	t := 0
	for i := 0; i < len(ar); i++ {
		n, er := strconv.Atoi(ar[i])
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Println("x:", x)
	fmt.Println("ar:", ar)
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!")

}
