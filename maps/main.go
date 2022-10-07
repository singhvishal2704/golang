package main

import "fmt"

func main(){
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["RB"] = "Ruby"
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}