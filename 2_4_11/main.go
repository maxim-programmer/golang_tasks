package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	fmt.Println("Следующее за числом", x, "число:", x+1)
	fmt.Println("Для числа", x, "предыдущее число:", x-1)
}
