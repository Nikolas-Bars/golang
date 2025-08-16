package main

import (
	"log"
	"net/http"
    "strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        cycles := 5
    if cyclesStr := r.URL.Query().Get("cycles"); cyclesStr != "" {
        // Преобразуем строку в число
        if c, err := strconv.Atoi(cyclesStr); err == nil {
            cycles = c
        }
    
    }
	
    Lissajous2(w, cycles)

	})
	// log.Fatal используется здесь для важной цели - обработки потенциальных ошибок, которые может вернуть
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
