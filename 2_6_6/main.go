package main

import (
	"fmt"
	"math"
)

func main() {
	var bit float64
	_, _ = fmt.Scan(&bit)
	fmt.Println(bit / math.Pow(2, 13))
}
