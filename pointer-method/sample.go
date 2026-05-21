package main

import "fmt"

type User struct {
	Name string
}

// method with pointer receiver
func (u *User) updateUserName(name string) {
	u.Name = name
}

func main() {
	user := User{Name: "Budi"}
	fmt.Println(user)
	user.updateUserName("Joko")
	fmt.Println(user)


	// create pointer with &{value}
	user2 := User{Name: "test"}
	fmt.Println(user2)
	user2.updateUserName("test2")
	fmt.Println(user2)
	
}
