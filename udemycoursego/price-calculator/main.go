package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.5, 1.1, 1.5}

	for _, tr := range taxRates {
		fm := filemanager.New("pricesssss.txt", fmt.Sprintf("result_%.0f.json", tr*100))
		//cmd := cmdmanager.New()
		priceJob := prices.New(fm, tr)
		err := priceJob.Proccess()

		if err != nil {
			fmt.Println("Could Not process JOB")
			fmt.Println(err)
		}
	}

}
