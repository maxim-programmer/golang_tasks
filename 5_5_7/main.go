package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	num := make([]int, n)
	cnt1 := 0
	cnt2 := 0
	cnt3 := 0
	cntcars := 0
	for i := range n {
		_, _ = fmt.Scan(&num[i])
		switch num[i] {
		case 1:
			cnt1++
		case 2:
			cnt2++
		case 3:
			cnt3++
		case 4:
			cntcars++
		}
	}
	if cnt1 >= cnt3 {
		cnt1 -= cnt3
	} else {
		cnt1 = 0
	}
	if cnt2%2 == 1 {
		if cnt1 >= 2 {
			cnt1 -= 2
		} else {
			cnt1 = 0
		}
	}
	if cnt1%4 != 0 {
		cntcars++
	}
	cntcars += cnt3 + cnt2/2 + cnt2%2 + cnt1/4
	fmt.Println(cntcars)
}
