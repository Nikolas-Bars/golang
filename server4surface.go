package main

import (
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            width := 600
            xrange := 30.0
            angle := math.Pi / 6
            xyscale := float64(width)/2/xrange
        if cyclesStr := r.URL.Query().Get("angle"); cyclesStr != "" {
            // Преобразуем строку в число
            if c, err := strconv.ParseFloat(cyclesStr, 4); err == nil {
                angle = c
            }
        }
        if cyclesStr := r.URL.Query().Get("xyscale"); cyclesStr != "" {
            // Преобразуем строку в число
            if x, err := strconv.ParseFloat(cyclesStr, 4); err == nil {
                xyscale = x
            }
        }
        w.Header().Set("Content-Type", "image/svg+xml")
        Surface(w, angle, xyscale)

	})
    log.Println("Сервер запущен на http://localhost:8000")
	// log.Fatal используется здесь для важной цели - обработки потенциальных ошибок, которые может вернуть
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
