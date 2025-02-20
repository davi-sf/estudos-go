package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.5, 1.1, 1.5}
	doneChans := make([]chan bool, len(taxRates))
	erroChans := make([]chan error, len(taxRates))
	for i, tr := range taxRates {
		doneChans[i] = make(chan bool)
		erroChans[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", tr*100))
		//cmd := cmdmanager.New()
		priceJob := prices.New(fm, tr)

		go priceJob.Proccess(doneChans[i], erroChans[i])

		//err := priceJob.Proccess()
		//if err != nil {
		//fmt.Println("Could Not process JOB")
		//fmt.Println(err)
		//}
	}

	for index := range taxRates {
		select {
		case err := <-erroChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("done")
		}

	}

}
