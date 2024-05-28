package main

import "fmt"

// if
// if else
// else if
// if bersarang

func main() {
	// age := 10

	// if age > 10 {
	// 	fmt.Println("Boleh bermain game")
	// } else {
	// 	fmt.Println("Belum boleh bermain game")
	// }

	score := 65
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println(grade)
}
