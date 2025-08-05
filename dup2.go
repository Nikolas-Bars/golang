package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Создаем карту counts, которая хранит:
	// ключ - строка из файла,
	// значение - карта (имя файла -> сколько раз эта строка встретилась в файле)
	counts := make(map[string]map[string]int)

	// Получаем список имен файлов из аргументов командной строки
	files := os.Args[1:]

	// Если файлов не передано (например, запустили просто go run dup.go)
	if len(files) == 0 {
		// Считаем строки из стандартного ввода (клавиатуры)
		countLines("stdin", os.Stdin, counts)
	} else {
		// Иначе перебираем каждый файл из аргументов
		for _, filename := range files {
			// Пытаемся открыть файл
			f, err := os.Open(filename)
			if err != nil {
				// Если ошибка (файл не открылся), выводим сообщение об ошибке и переходим к следующему файлу
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// Если файл открылся успешно, считываем строки из файла и обновляем counts
			countLines(filename, f, counts)
			// Закрываем файл после чтения
			f.Close()
		}
	}

	// Проходим по всем строкам и для каждой проверяем, в каких файлах она встречалась больше одного раза
	for line, filemap := range counts {
		for filename, n := range filemap {
			if n > 1 {
				// Выводим строку, сколько раз она встретилась и в каком файле
				fmt.Printf("%q встречается %d раз в файле: %s\n", line, n, filename)
			}
		}
	}
	//map[книга:map[text1.txt:2] кот:map[text1.txt:2] лошадь:map[text1.txt:1] рука:map[text2.txt:2]]
	fmt.Println(counts)
}

// Функция считывает строки из переданного файла и обновляет карту counts
// filename - имя файла (нужно для записи в counts)
// f - указатель на открытый файл
// counts - карта для хранения строк и их количества в файлах
func countLines(filename string, f *os.File, counts map[string]map[string]int) {
	// Создаем сканер, который будет читать файл построчно
	input := bufio.NewScanner(f)

	// Пока есть строки для чтения
	for input.Scan() {
		line := input.Text() // читаем строку из файла

		// Если для этой строки еще нет вложенной карты (имя файла -> количество),
		// то создаем пустую карту
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		// Увеличиваем счетчик для этой строки и файла на 1
		counts[line][filename]++
	}
}
