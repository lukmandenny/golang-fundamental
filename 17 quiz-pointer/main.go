package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

// AddGame

func (game *Gamer) AddGame(namagame string) {
	game.Games = append(game.Games, namagame)
}

func main() {
	game := Gamer{Name: "MMORPG"}

	game.AddGame("Moonlight Sculptor: Dark Gamer")
	game.AddGame("Ragnarok Origin Global")

	for _, game := range game.Games {
		fmt.Println(game)
	}

	fmt.Println(game)

}
