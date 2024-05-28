package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) graduate() {
	student.Name = student.Name + ", S.Tr.Kom"
}

func main() {
	student := Student{1, "Lukman Denny", 3.04}
	fmt.Println(student.Name)
	student.graduate()
	fmt.Println(student.Name)
}
