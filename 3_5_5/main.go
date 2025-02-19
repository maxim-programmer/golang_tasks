package main

import "fmt"

func main() {
	var c int
	_, _ = fmt.Scan(&c)
	if c <= 3 {
		fmt.Println("начинающий")
	} else if c >= 4 && c <= 7 {
		fmt.Println("младший разработчик")
	} else if c >= 8 && c <= 15 {
		fmt.Println("средний разработчик")
	} else {
		fmt.Println("старший разработчик")
	}
}
