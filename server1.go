package main

import (
	"fmt"
	"log"
	"net/http"
)

// Главная функция программы
func main() {
	// Регистрируем обработчик (handler) для всех запросов по пути "/"
	// http.HandleFunc принимает два аргумента:
	// 1. Путь, для которого регистрируется обработчик ("/" - корневой путь)
	// 2. Функцию-обработчик (в данном случае handler)
	http.HandleFunc("/", handler)

	// Запускаем HTTP-сервер и логируем возможные ошибки
	// http.ListenAndServe принимает:
	// 1. Адрес сервера в формате "host:port" ("localhost:8000")
	// 2. Обработчик маршрутов (nil означает использование DefaultServeMux)
	// log.Fatal завершает программу с кодом 1 и выводит ошибку, если сервер не смог запуститься
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Функция-обработчик HTTP-запросов
// Принимает два параметра:
// 1. w http.ResponseWriter - интерфейс для формирования HTTP-ответа
// 2. r *http.Request - указатель на структуру с информацией о HTTP-запросе
func handler(w http.ResponseWriter, r *http.Request) {
	// Форматируем и отправляем клиенту строку с путем из URL
	// fmt.Fprintf записывает данные не в стандартный вывод, а в ResponseWriter
	// %q - спецификатор формата, который добавляет кавычки и экранирует спецсимволы
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}