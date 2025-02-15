package main

import (
	"errors"
	"fmt"
	"os"
)

func calculation(revenueValue float64, expenses float64, taxRate float64) (float64, float64, float64) {

	ebt := revenueValue - expenses
	profit := ebt - (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio

}

func valuesInput(infoText string) (float64, error) {
	var value float64

	for {
		fmt.Print(infoText)
		fmt.Scan(&value)

		if value > 0 {
			return value, nil
		}

		fmt.Println(errors.New("ERRO! Positive numbers only"))
	}
}

func saveValuesFile(ebt, profit, ratio float64) {
	balanceText := fmt.Sprintf("Ebt: %.2f \nProfit: %.2f \nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile("results", []byte(balanceText), 0644)
}

func main() {

	revenueValue, _ := valuesInput("Revenue value: ")
	expenses, _ := valuesInput("Expenses value: ")
	taxRate, _ := valuesInput("Tax Rate: ")

	ebt, profit, ratio := calculation(revenueValue, expenses, taxRate)
	saveValuesFile(ebt, profit, ratio)

	fmt.Printf("Ebt: %.2f", ebt)
	fmt.Printf("\nProfit: %.2f", profit)
	fmt.Printf("\nRatio: %.2f", ratio)

}
