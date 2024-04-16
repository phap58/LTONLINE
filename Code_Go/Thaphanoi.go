package main

import (
	"fmt"
)

func towerOfHanoi(n int, from string, to string, aux string, d *int) {
	if n == 1 {
		*d++
		fmt.Printf("Buoc%d Chuyen dia 1 tu %s sang %s\n", *d, from, to)
		return
	}
	towerOfHanoi(n-1, from, aux, to, d)
	*d++
	fmt.Printf("Buoc%d Chuyen dia %d tu %s sang %s\n", *d, n, from, to)
	towerOfHanoi(n-1, aux, to, from, d)
}

func main() {
	var n int
	fmt.Scan(&n)
	d := 0
	towerOfHanoi(n, "A", "B", "C", &d)
}
