package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

type Asal struct {
	Angka int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}
func (PersegiPanjang PersegiPanjang) HitungLuas() int {
	return PersegiPanjang.Panjang * PersegiPanjang.Lebar
}
func (asal Asal) HitungLuas() int {
	return 1011
}

func main() {
	persegi := Persegi{8}
	result := SeberapaLuas(persegi)
	fmt.Println("Luas dari persegi dengan sisi", persegi.Sisi, "adalah:", result)

	persegiPanjang := PersegiPanjang{8, 9}
	result = SeberapaLuas(persegiPanjang)
	fmt.Println("Luas dari persegi panjang dengan panjang", persegiPanjang.Panjang, "dan lebar", persegiPanjang.Lebar, "adalah:", result)

	asal := Asal{9}
	result = SeberapaLuas(asal)
	fmt.Println("Luas dari asal dengan angka", asal.Angka, "adalah:", result)

}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
