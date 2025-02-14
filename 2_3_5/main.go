package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	separator := scanner.Text()
	_ = scanner.Scan()
	input1 := scanner.Text()
	_ = scanner.Scan()
	input2 := scanner.Text()
	_ = scanner.Scan()
	input3 := scanner.Text()
	fmt.Println(input1 + separator + input2 + separator + input3)
}
