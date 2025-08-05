package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	// range - переберет данные массива в указанном диапазоне
	for i, arg := range os.Args[0:] {
		// arg - это будет элемент под соответствуюшим индексом
		s += sep + arg
		sep = " "
		fmt.Printf("%#v \n", i)
	}

	fmt.Printf("%#v %#v \n", s, s)
	fmt.Println(strings.Join(os.Args[0:], " "))
}

// при команде go run main.go --varsion pop jop -fuck
// мы получим string "--varsion pop jop -fuck"
