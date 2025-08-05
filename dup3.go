package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Создаём словарь для подсчёта количества одинаковых строк
	counts := make(map[string]int)

	// Перебираем имена всех файлов, переданных при запуске программы
	for _, filename := range os.Args[1:] {
		// Читаем весь файл как массив байтов
		data, err := os.ReadFile(filename)
		if err != nil {
			// Если файл не открылся — выводим ошибку и переходим к следующему - (os.Stderr - стандартный поток ошибок)
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}

		// Выводим содержимое файла как текст (для отладки или проверки)
		//fmt.Println(string(data))

		// Преобразуем байты в строку, потом разбиваем на строки по символу новой строки (\n)
		lines := strings.Split(string(data), "\n")
		// Считаем, сколько раз каждая строка встретилась
		for _, line := range lines {
			counts[line]++
		}
	}

	// Выводим строки, которые встретились более одного раза
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
