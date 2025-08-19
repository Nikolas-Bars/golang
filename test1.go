package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	cwd, err := os.Getwd() // Примечание: неверно!
	//var i byte
	//i = 255
	f := 260
	z := uint8(f)
	fmt.Printf("%#v\n %[2]d\n %[2]f\n", z, f) // 0
	if err != nil {
		log.Fatalf("Ошибка os.Getwd: %v", err)
	}
	log.Printf("Рабочий каталог = %s", cwd)
}
func main() {

}