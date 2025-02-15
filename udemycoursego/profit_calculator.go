package main

import "fmt"

func calculation(revenueValue float64, expenses float64, taxRate float64) (float64, float64, float64) {

	ebt := revenueValue - expenses
	profit := ebt - (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio

}

func valuesInput(infoText string) float64 {

	var value float64

	fmt.Print(infoText)
	fmt.Scan(&value)

	return value

}

func main() {

	revenueValue := valuesInput("Revenue value: ")
	expenses := valuesInput("Expenses value: ")
	taxRate := valuesInput("Tax Rate: ")

	ebt, profit, ratio := calculation(revenueValue, expenses, taxRate)

	fmt.Printf("Ebt: %.2f", ebt)
	fmt.Printf("Profit: %.2f", profit)
	fmt.Printf("Ratio: %.2f", ratio)

}
