package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := -25; i <= 15; i++ {
		i := float64(i)
		sum += y(i)
	}
	fmt.Println(sum)
}

func y(x float64) int {
	var y float64
	if x < -5 {
		y = 2*math.Abs(x) - 1
	} else if -5 <= x && x <= 5 {
		y = x
	} else {
		y = 2 * x
	}
	return int(y)
}
