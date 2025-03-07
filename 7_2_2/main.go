package main

import "fmt"

func main() {
	var n, m float64
	_, _ = fmt.Scan(&n, &m)
	fmt.Println(sr_ar(n) + sr_ar(m))
}

func sr_ar(x float64) float64 {
	var sum float64 = 0
	for i := 1; i <= int(x); i++ {
		sum += float64(i)
	}
	return sum / x
}
