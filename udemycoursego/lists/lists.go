package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	//1)
	hobbies := [3]string{"guitaring", "cooking", "gaming"}
	fmt.Println(hobbies)

	//2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3)
	hobbiesList := hobbies[:2]
	fmt.Println(hobbiesList)

	// 4)
	courseGoals := []string{"Learning GO", "To be ready for the internship"}
	courseGoals = append(courseGoals, "Learn everything I can ")
	fmt.Println(courseGoals)

	//5)
	productList := []Product{
		{
			"first-product",
			1,
			47.90,
		},
		{
			"second-product",
			2,
			10.99,
		},
	}

	fmt.Println(productList)

	newProduct := Product{
		"third-product",
		3,
		43.50,
	}

	productList = append(productList, newProduct)
	fmt.Println(productList)
}
