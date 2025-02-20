package main

import "fmt"

func main() {
	var n int
	cnt_otr := 0
	cnt_pol := 0
	_, _ = fmt.Scan(&n)
	for n != 0 {
		if n > 0 {
			cnt_pol++
		} else {
			cnt_otr++
		}
		_, _ = fmt.Scan(&n)
	}
	fmt.Println(cnt_pol - cnt_otr)
}
