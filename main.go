package main

import "fmt"

func main() {

	usdToEur, usdToRub := outputUser()

	conv := usdToRub / usdToEur

	fmt.Print(conv)

}

func outputUser() (float64, float64) {

	var usdToEur float64

	var usdToRub float64

	fmt.Scan(&usdToEur)

	fmt.Scan(&usdToRub)

	return usdToEur, usdToRub
}

func calculatCurrency(usd float64, rub float64, eur float64) {

}
