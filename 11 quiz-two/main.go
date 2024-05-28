package main

import "fmt"

func main() {
	// hitung rata-rata
	score := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	totalScore := 0

	for _, value := range score {
		totalScore = totalScore + value
	}
	length := len(score)
	average := float64(totalScore) / float64(length)
	fmt.Println("Rata-rata = ", average)

	// score := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	goodScores := []int{}

	for _, value := range score {
		if value >= 90 {
			goodScores = append(goodScores, value)
		}
	}
	fmt.Println("Nilai baik:", goodScores)
}
