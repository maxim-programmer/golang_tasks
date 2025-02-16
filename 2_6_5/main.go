package main

import "fmt"

func main() {
	var x float64
	_, _ = fmt.Scan(&x)
	fmt.Println(x - float64(int(x)))
}
