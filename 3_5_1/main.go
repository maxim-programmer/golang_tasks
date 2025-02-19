package main

import "fmt"

func main() {
	var age int
	var pol string
	_, _ = fmt.Scan(&age, &pol)
	if pol == "m" && age >= 12 && age <= 18 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
