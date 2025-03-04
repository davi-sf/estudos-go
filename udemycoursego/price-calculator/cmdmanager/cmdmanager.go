package cmdmanager

import (
	"fmt"
)

type CMDManager struct {
}

func (fm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Enter prices. Confirm every price with enter")

	var prices []string
	for {
		var price string
		fmt.Println("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (fm CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
