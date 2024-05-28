package main

import "fmt"

func main() {
	// numberA := 5
	// numberB := &numberA //referencing dengan tanda & untuk mengambil alamat memori dari numberA
	// fmt.Println(numberA)
	// fmt.Println(*numberB) //dereferencing menggunakan tanda *, untuk menampilkan variabel yang disimpan
	// *numberB = 10
	// fmt.Println(numberA)

	// var numberA int = 5
	// var numberB *int = &numberA
	// fmt.Println(numberA)
	// fmt.Println(numberB)

	// numberA = 20
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	number := 5
	fmt.Println("Alamat memori awal:", &number)
	fmt.Println("Nilai awal:", number)
	change(&number, 100)
	fmt.Println("Alamat memori baru:", &number)
	fmt.Println("Nilai baru:", number)

}

func change(old *int, new int) {
	*old = new
	fmt.Println("alamat memori di proses:", old)
	fmt.Println("Nilai di proses:", *old)
}
