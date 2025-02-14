package main

import "fmt"

func main() {
	var minutes int
	_, _ = fmt.Scan(&minutes)
	fmt.Println(minutes, "мин - это", minutes/60, "час", minutes%60, "минут.")
}
