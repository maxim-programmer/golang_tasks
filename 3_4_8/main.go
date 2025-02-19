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
	if x1 == x2 || y1 == y2 {
		fmt.Println("YES")
	} else if math.Abs(x2-x1) == math.Abs(y2-y1) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
