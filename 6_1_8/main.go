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
		if s[i] != 32 {
			fmt.Print(string(s[i]))
		} else if s[i] == 32 && s[i+1] != 32 {
			fmt.Print(string(s[i]))
		}
	}
}
