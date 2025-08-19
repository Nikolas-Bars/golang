package main

// pc[i] - количество единичных битов в i.

var pc [256]byte

func init() {
		for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
// Здесь x — 64-битное число. Его нужно разбить на 8 байтов (по 8 бит = 1 байт).
func PopCount(x uint64) int {
	return int(
		//x >> N сдвигает биты числа x вправо на N позиций.
		// Каждый сдвиг вправо на 1 позицию делит число на 2 (без остатка).
		pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	PopCount(45666666666)
}