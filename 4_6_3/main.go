package main

import "fmt"

func main() {
	var n float64
	var sum float64 = 0
	var cnt float64 = 0
	_, _ = fmt.Scan(&n)
	for n != 0 {
		sum += n
		cnt++
		_, _ = fmt.Scan(&n)
	}
	fmt.Println(sum / cnt)
}
