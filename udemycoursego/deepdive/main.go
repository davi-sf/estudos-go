package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	double := transformNumber(&nums, triple)
	fmt.Println(double)

}

type transformFn func(int) int

func transformNumber(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}

	return dNumbers

}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}
