package main

import "fmt"

func main() {
	websites := map[string]string{
		"google": "https://www.google.com",
		"github": "https://github.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Github"])
	websites["linkedin"] = "https://www.linkedin.com"
	delete(websites, "google")
	fmt.Println(websites)
}
