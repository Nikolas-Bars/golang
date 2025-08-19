package main

import (
	"fmt"
	"os"
	"strconv"
	"golang/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v\n", tempconv.CToF(tempconv.Celsius(x)))
		fmt.Printf("%v\n", tempconv.FToC(tempconv.Fahrenheit(x)))
	}
}