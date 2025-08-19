package main

import (
	"log"
	"os"
)

var cwd string

func init() {
	cwd, err := os.Getwd() // Примечание: неверно!
	if err != nil {
		log.Fatalf("Ошибка os.Getwd: %v", err)
	}
	log.Printf("Рабочий каталог = %s", cwd)
}
func main() {

}