package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// создаем таблицу ключ (string) значение(int) (похож на объект js)
	counts := make(map[string]int)
	// добавляем сканер строк в терминале
	input := bufio.NewScanner(os.Stdin) // os.Stdin - ввод с клавиатуры
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%v %v \n", n, line)
		}
	}
	//map[string]int map[string]int{"bob":1, "bob ":1, "cat":1, "dog":1, "nat":1, "team":2}
	fmt.Printf("%T %v \n", counts, counts)
}
