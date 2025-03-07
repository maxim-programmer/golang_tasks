package main

import "fmt"

func main() {
	var n1, n2 int
	_, _ = fmt.Scan(&n1, &n2)
	if yes(n1)*yes(n2) == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}

func yes(number int) int {
	if (number/100000 + number/10000%10 + number/1000%10) == (number/100%10 + number/10%10 + number%10) {
		return 1
	}
	return 0
}
