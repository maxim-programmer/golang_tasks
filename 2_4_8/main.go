package main

import "fmt"

func main() {
	var phone_price int
	var case_price int
	var charger_price int
	var headphones_price int
	_, _ = fmt.Scan(&phone_price, &case_price, &charger_price, &headphones_price)
	fmt.Println(3 * (phone_price + case_price + charger_price + headphones_price))
}
