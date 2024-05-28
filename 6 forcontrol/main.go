package main

import "fmt"

func main() {
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Saya sedang belajar Go", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Saya sedang belajar Go", i)
	// 	i++
	// }

	title := "Golang the best language"
	// for index, letter := range title {
	// 	fmt.Println("Index:", index, "Letter:", string(letter))
	// }
	// for _, letter := range title {
	// 	fmt.Println("Letter:", string(letter))
	// }

	// QUIZ
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("Index:", index, "Letter:", string(letter))
		}
	}

	fmt.Println("Vocal")
	for index, letter := range title {
		let := string(letter)
		switch let {
		case "a", "i", "u", "e", "o":
			fmt.Println("Index:", index, "Letter:", let)
		}

	}
}
