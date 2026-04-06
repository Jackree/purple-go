package main

import "fmt"

func main() {
	const USDEUR = 0.86
	const USDRUB = 80.45
	EURRUB := USDRUB / USDEUR
	fmt.Print("EURRUB: ", EURRUB)
}

func getUserInput() string {
	var input string
	fmt.Scan(&input)
	
	return input
}

func calculate(num float64, fromValue string, toValue string) float64 {
	
}