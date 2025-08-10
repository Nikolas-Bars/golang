package main

import (
	"fmt" // пакет для форматированного вывода
	"io"  // пакет для работы с потоками ввода/вывода
	"net/http" // пакет для выполнения HTTP-запросов
	"os"    
	"time"   // пакет для работы с аргументами, stderr и выходом
)

func main() {
    // Засекаем общее время выполнения всей программы
    start := time.Now()

    // Создаём канал (chanel), в который горутины будут отправлять строки с результатами
    ch := make(chan string)

    // Перебираем все URL'ы, переданные в командной строке (os.Args[1:] — без имени программы)
    for _, url := range os.Args[1:] {
        // Запускаем функцию fetch в отдельной горутине для каждого URL
        // Это позволяет обрабатывать несколько URL одновременно (параллельно)
        go fetch(url, ch)
    }

    // Получаем и печатаем результаты от всех горутин
    // Количество получений из канала = количеству URL'ов
    for range os.Args[1:] {
        // Получаем сообщение из канала (результат одной fetch) и выводим
        fmt.Println(<-ch)
    }

    // Выводим общее время выполнения всей программы
    // time.Since(start) — сколько времени прошло с начала
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
// fetch выполняет HTTP-запрос к указанному URL, измеряет длительность,
// читает весь ответ и отправляет результат в канал ch -chan<- string означает, 
// что функция только отправляет в канал (не читает из него)
func fetch(url string, ch chan<- string) {
    // Засекаем время начала обработки текущего URL
    start := time.Now()

    // Выполняем HTTP GET-запрос по указанному URL
    resp, err := http.Get(url)

    // Если произошла ошибка (например, сайт недоступен), отправляем сообщение об ошибке в канал и выходим
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }

    // Читаем всё тело ответа, но не сохраняем его — отправляем в io.Discard
    // Это значит, что мы просто измеряем, сколько данных передано (в байтах),
    // не загружая их в память
    nbytes, err := io.Copy(io.Discard, resp.Body)

    // Закрываем тело ответа, чтобы освободить ресурсы
    resp.Body.Close()

    // Если возникла ошибка при чтении данных — сообщаем в канал и выходим
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }

    // Вычисляем, сколько времени заняла обработка текущего URL
    secs := time.Since(start).Seconds()

    // Формируем строку с результатом:
    // - время загрузки в секундах,
    // - количество байт,
    // - URL
    // и отправляем её в канал
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}


// ПРИМЕ С ЗАПИСЬЮ В ФАЙЛ

// func fetch(url string, ch chan<- string) {
// 	start := time.Now()
// 	resp, err := http.Get(url)

// 	if err != nil {
// 		ch <- fmt.Sprintf("error fetching %s: %v", url, err)
// 		return
// 	}

// 	nbytes, err := io.Copy(io.Discard, resp.Body)
// 	resp.Body.Close()

// 	if err != nil {
// 		ch <- fmt.Sprintf("error reading %s: %v", url, err)
// 		return
// 	}

// 	secs := time.Since(start).Seconds()
// 	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
// }

// func main() {
// 	// Открываем файл для записи (если не существует — создаём, если существует — перезаписываем)
// 	file, err := os.Create("fetchall.log")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error creating file: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer file.Close() // Закрываем файл при завершении программы

// 	start := time.Now()
// 	ch := make(chan string)

// 	// Запускаем горутины для каждого URL
// 	for _, url := range os.Args[1:] {
// 		go fetch(url, ch)
// 	}

// 	// Собираем результаты
// 	for range os.Args[1:] {
// 		result := <-ch
// 		fmt.Println(result)                  // Вывод в консоль (опционально)
// 		fmt.Fprintln(file, result)           // Запись в файл
// 	}

// 	// Записываем общее время выполнения
// 	totalTime := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
// 	fmt.Print(totalTime)
// 	fmt.Fprint(file, totalTime)
// }