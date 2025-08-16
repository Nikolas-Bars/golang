package main

import (
	"fmt"
	"os"
)
const boilingF = 212.0
func main() {
	var f = boilingF
	var c = (f - 32) *5/9
	fmt.Printf("Температура кипения = %g°F или %g°C\n", f, c)
	// Вывод:
	// Температура кипения = 212°F или 100°C

	in, err := os.Open(infile)
	l, err := os.Create(outfile)
}