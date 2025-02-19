package main

import "fmt"

func main() {
	var x float64
	_, _ = fmt.Scan(&x)
	if x == 0 {
		fmt.Println("Обратного числа не существует")
	} else {
		fmt.Println(1 / x)
	}
}
