package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	s1 := scanner.Text()
	_ = scanner.Scan()
	s2 := scanner.Text()
	if string([]rune(s1)[0]) == string([]rune(s2)[len(s2)-1]) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
