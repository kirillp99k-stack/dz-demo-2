package main

import "fmt"

func main() {
	const usdInEuro = 0.8541
	const usdInRub = 83.61
	var text string = " EUR/RUB = KOLYAS"
	fmt.Print(text, usdInRub/usdInEuro)

}
