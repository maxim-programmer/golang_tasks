package main

import (
	"fmt"
	"math"
)

func main() {
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64
	_, _ = fmt.Scan(&x1, &y1, &x2, &y2)
	fmt.Println(math.Abs(x1-x2) + math.Abs(y1-y2))
}
