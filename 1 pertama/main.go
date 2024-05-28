package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Go")
	// sentence := TestAja("Denny", 2024)
	// fmt.Println(sentence)
	var a = 9
	var b = 8
	Add := calculation.Add(a, b)
	Multiply := calculation.Multiply(a, b)
	fmt.Println("Hasil penjumlahan dari", a, "dan", b, "adalah", Add)
	fmt.Println("Hasil perkalian dari", a, "dan", b, "adalah", Multiply)
}
