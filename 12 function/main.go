package main

import "fmt"

func main() {
	sentence := printMyResult("Saya sedang")
	fmt.Println(sentence)

	result := add(8, 9)
	fmt.Println(result)

	_, x, _ := calculate(8, 9)
	// fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", x)
	// fmt.Println("Jumlah:", jumlah)
	sentence2 := printMyType("Mantap!")
	fmt.Println(sentence2)

}

func printMyResult(sentence string) string {
	newSentence := sentence + " belajar Go"
	return newSentence
}

// pre-defined return value
// variable newSentence yang digunakan sebagai return value dalam fungsi bisa langsung di definisikan di awal
// sehingga return secara otomatis ke newSentence
func printMyType(sentence string) (newSentence string) {
	newSentence = sentence
	return
}

func add(number, numberTwo int) int {
	return number + numberTwo
}

// bisa multiple proses dan return dengan sekali input
func calculate(panjang int, lebar int) (int, int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	jumlah := panjang + lebar
	return luas, keliling, jumlah
}

// 1. input
// 2. proses
// 3. output
