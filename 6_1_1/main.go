package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	s := scanner.Text()
	cnt := 1
	for i := range len(s) {
		if string(s[i]) == " " {
			cnt++
		}
	}
	fmt.Println(cnt)
}
