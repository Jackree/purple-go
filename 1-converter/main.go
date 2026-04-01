package main

import "fmt"

func main() {
	const USDEUR = 0.86
	const USDRUB = 80.45
	EURRUB := USDRUB / USDEUR
	fmt.Print("EURRUB: ", EURRUB)
}