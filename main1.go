// ТУТ НАШ ПАКЕТ MAIN!!!
package main

import (
	"fmt"
	"golang/tempconv"
)

func main() {
	fmt.Printf("%v\n", tempconv.CToF(30))
	fmt.Printf("%v\n", tempconv.FToC(30))
	fmt.Printf("%v\n", tempconv.KToC(100))
}
