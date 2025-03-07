package main

import "fmt"

func main() {
	var f, i, o string
	_, _ = fmt.Scan(&f, &i, &o)
	fmt.Println(f, string([]rune(i)[0])+"."+string([]rune(o)[0])+".")
}
