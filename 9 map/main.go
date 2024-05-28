package main

import "fmt"

func main() {
	// myMap := map[string]int{}

	// myMap["dart"] = 8
	// myMap["go"] = 9
	// myMap["java"] = 7

	// fmt.Println(myMap)

	myMap := map[string]int{
		"dart": 8,
		"go":   9,
		"java": 7,
	}

	// for key, value := range myMap {
	// 	fmt.Println("Key:", key, "Value:", value)
	// }

	// fmt.Println("=================================")

	// delete(myMap, "java")

	// for key, value := range myMap {
	// 	fmt.Println("Key:", key, "Value:", value)
	// }

	value, isAvailable := myMap["pyhton"]
	fmt.Println(value)
	fmt.Println(isAvailable)
}
