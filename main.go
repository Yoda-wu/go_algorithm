package main

import (
	"fmt"
)

const (
	A = 2
	B = 3
)

func main() {
	var f [A][B]int
	f[1][1] = 0
	fmt.Println(f[1][1])

}
