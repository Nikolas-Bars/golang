package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
    mu    sync.Mutex // Мьютекс для синхронизации доступа к счетчику
    count int        // Счетчик запросов
)

// Главная функция программы
func main() {
    // Регистрируем обработчик для корневого пути "/"
    // Будет вызываться для всех запросов, кроме "/count"
    http.HandleFunc("/", handler)
    
    // Регистрируем обработчик для пути "/count"
    // Будет отображать текущее количество запросов
    http.HandleFunc("/count", counter)
    
    // Запускаем HTTP-сервер на localhost порт 8000
    // log.Fatal завершит программу с выводом ошибки, если сервер не запустится
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработчик корневого пути "/"
func handler(w http.ResponseWriter, r *http.Request) {
    // Блокируем мьютекс для безопасного увеличения счетчика
    // Это предотвращает состояние гонки (race condition) при конкурентных запросах
    mu.Lock()
    count++ // Увеличиваем счетчик запросов
    mu.Unlock() // Освобождаем мьютекс
    
    // Отправляем клиенту ответ с путем запроса
    // %q форматирует строку с кавычками и экранированием спецсимволов
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Обработчик пути "/count"
func counter(w http.ResponseWriter, r *http.Request) {
    // Блокируем мьютекс для безопасного чтения счетчика
    mu.Lock()
    // Отправляем клиенту текущее значение счетчика
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock() // Освобождаем мьютекс
}