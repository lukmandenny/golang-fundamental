package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println("Total:", total)

	result, err := calculate(10, 2, "l")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Hasil Operasi:", result)

}

func sum(numbers []int) (total int) {
	for _, number := range numbers {
		fmt.Println(total, "+", number)
		total = total + number
	}
	return
}

func calculate(number, numberTwo int, operasi string) (result int, err error) {
	switch operasi {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		err = errors.New("JENIS OPERASI TIDAK DIKETAHUI")
	}
	return
}
