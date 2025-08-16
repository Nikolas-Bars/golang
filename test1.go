package main

import (
	"fmt"
	"os"
)

var y *int

func main() {
	f()
	fib(6)
}

func f() {
	i := 5
	y = &i
	printY()
}

func printY() {
	fmt.Printf("%v\n", *y)
}

func g() {
	i := new(int)
	*i = 1
}

func fib (n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		// x = 1, y = 0 + 1
		// x = 1, y = 1 + 1
		// x = 2, y = 2 + 1
		// x = 3, y = 3 + 2
		// x = 5, y = 5 + 3
		// x = 8, y = 8 + 5
		x, y = y, x + y
	}
	fmt.Printf("%v\n", x)
	os.Open()
}