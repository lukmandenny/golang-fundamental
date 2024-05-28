package main

import "fmt"

func main() {
	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "JavaScript"
	// languages[3] = "C#"
	// languages[4] = "Python"

	languages := [...]string{
		"Go", "Dart", "Java", "Python", "Php",
	}

	// fmt.Println(languages)
	// lenght := len(languages)
	// fmt.Println(lenght)

	for index, lang := range languages {
		fmt.Println("Index:", index, "Language:", lang)
	}

}
