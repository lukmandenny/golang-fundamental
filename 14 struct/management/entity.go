package management

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// function method
func (user User) Display() string {
	return fmt.Sprintf("Nama : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	// fmt.Printf: Mengformat string dan langsung mencetaknya ke standar output.
	// fmt.Sprintf: Mengformat string dan mengembalikan string yang telah diformat tanpa mencetaknya.
}

func (group Group) DisplayGroup() {
	fmt.Printf("Nama Group: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Jumlah Anggota: %d", len(group.Users))
	fmt.Println("")
	fmt.Println("Nama Anggota:")
	for i, user := range group.Users {
		fmt.Println(i+1, ".", user.FirstName, user.LastName)
	}
	fmt.Println("Detail Anggota:")
	for i, user := range group.Users {
		result := user.Display()
		fmt.Println(i+1, ".", result)
	}
}
