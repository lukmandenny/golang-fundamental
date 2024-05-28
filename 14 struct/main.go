package main

import "struct/management"

func main() {
	user := management.User{}
	user.ID = 1
	user.FirstName = "Lukman"
	user.LastName = "Denny"
	user.Email = "lukmandenny@gmail.com"
	user.IsActive = true

	user2 := management.User{
		ID:        2,
		FirstName: "Baja",
		LastName:  "Hitam",
		Email:     "bajahitam@gmail.com",
		IsActive:  true,
	}

	// result := displayUser(user)
	// result2 := displayUser(user2)
	// fmt.Println(result)
	// fmt.Println(result2)

	users := []management.User{user, user2}

	group := management.Group{
		Name:        "Pahlawan Bertopeng",
		Admin:       user,
		Users:       users,
		IsAvailable: true,
	}

	// displayGroup(group)
	group.DisplayGroup()

}

// Struct bisa jadi parameter dari sebuah function
// func displayUser(user User) string {
// 	result := fmt.Sprintf("Nama : %s %s, Email : %s", user.FirstName, user.LastName, user.email)
// 	return result
// }

// func displayGroup(group Group) {
// 	fmt.Printf("Nama Group: %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Jumlah Anggota: %d", len(group.Users))
// 	fmt.Println("")
// 	fmt.Println("Nama Anggota:")
// 	for i, user := range group.Users {
// 		fmt.Println(i+1, ".", user.FirstName, user.LastName)
// 	}
// 	fmt.Println("Detail Anggota:")
// 	for i, user := range group.Users {
// 		result := user.display()
// 		fmt.Println(i+1, ".", result)
// 	}
// }
