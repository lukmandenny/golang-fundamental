package main

import "fmt"

func main() {
	// var gamingConsoles []string

	// gamingConsoles = append(gamingConsoles, "Xbox series s")

	// fmt.Println(len(gamingConsoles))

	gamingConsoles := []string{
		"Xbox Series S",
		"Nintendo Switch",
		"Steamdeck",
	}

	gamingConsoles = append(gamingConsoles, "Playstation 5")

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}

	fmt.Println(len(gamingConsoles))

}
