package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	_, _ = fmt.Scan(&a, &b)
	fmt.Println(math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2)))
}
