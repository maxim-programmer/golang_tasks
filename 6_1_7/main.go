package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	if (s[0] >= 65 && s[0] <= 90) || (s[0] >= 97 && s[0] <= 122) || s[0] == 95 {
		flag := false
		for i := range len(s) {
			if (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) || s[i] == 95 || (s[i] >= 48 && s[i] <= 57) {
				flag = true
			} else {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		fmt.Println("NO")
	}
}
