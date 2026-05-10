package main

import "fmt"

const (
	USDEUR = 0.86
	USDRUB = 80.45
)

func main() {
	fmt.Println("=== Конвертер валют ===")

	fromCurrency := getCurrencyInput("Введите исходную валюту (USD, EUR, RUB): ")

	amount := getNumberInput("Введите сумму: ")

	toCurrency := getCurrencyInput("Введите целевую валюту (USD, EUR, RUB): ")

	result := calculate(amount, fromCurrency, toCurrency)

	fmt.Printf(
		"\nРезультат: %.2f %s = %.2f %s\n",
		amount,
		fromCurrency,
		result,
		toCurrency,
	)
}

func getCurrencyInput(message string) string {
	for {
		var currency string

		fmt.Print(message)
		fmt.Scan(&currency)

		if currency == "USD" || currency == "EUR" || currency == "RUB" {
			return currency
		}

		fmt.Println("Ошибка! Доступны только: USD, EUR, RUB")
	}
}

func getNumberInput(message string) float64 {
	for {
		var number float64

		fmt.Print(message)

		_, err := fmt.Scan(&number)

		if err == nil {
			return number
		}

		fmt.Println("Ошибка! Введите корректное число.")
	}
}

func calculate(amount float64, from string, to string) float64 {
	if from == to {
		return amount
	}

	var usdAmount float64

	switch from {
	case "USD":
		usdAmount = amount

	case "EUR":
		usdAmount = amount / USDEUR

	case "RUB":
		usdAmount = amount / USDRUB
	}

	switch to {
	case "USD":
		return usdAmount

	case "EUR":
		return usdAmount * USDEUR

	case "RUB":
		return usdAmount * USDRUB
	}

	return 0
}
