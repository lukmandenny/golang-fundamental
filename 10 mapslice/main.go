package main

import "fmt"

func main() {
	students := []map[string]string{}

	students = append(students, map[string]string{
		"nama":  "Denny",
		"kelas": "Satu",
	})

	students = append(students, map[string]string{
		"nama":  "Dendut",
		"kelas": "Dua",
	})

	fmt.Println(students)
	fmt.Println("==============")

	for _, student := range students {
		fmt.Println(student["nama"])
	}
}
