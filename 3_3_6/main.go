package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64
	_, _ = fmt.Scan(&a, &b, &c)
	d := math.Sqrt(math.Pow(b, 2) - 4*a*c)
	if d > 0 {
		x1 := (-b - d) / (2 * a)
		x2 := (-b + d) / (2 * a)
		if x1 < x2 {
			fmt.Println(x1)
			fmt.Println(x2)
		} else {
			fmt.Println(x2)
			fmt.Println(x1)
		}
	} else {
		if d == 0 {
			fmt.Println((-b + d) / (2 * a))
		}
	}
}
