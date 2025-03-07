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
	for i := range len(s) {
		if 48 <= s[i] && s[i] <= 57 {
			fmt.Print(string(s[i]), " ")
		}
	}
}
