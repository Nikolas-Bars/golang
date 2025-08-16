package main

import (
	"flag"
	"fmt"
	"strings"
	// "os"
	// "strconv"
)
// flag.Bool создает новую переменную-флаг типа bool
var n = flag.Bool("n", false, "пропуск символа новой строки")
var sep = flag.String("s", ", ", "Разделитель")

//go run test.go -s ": " ц ц в в - результат ц: ц: в: в
func main() {
	// Проверяет все аргументы, которые вы передали программе при запуске Распознает флаги (параметры, начинающиеся с -)
	flag.Parse()
	// Аргументы, не являющиеся флагами, доступны через f l a g .A r g s ( ) как срез строк
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}