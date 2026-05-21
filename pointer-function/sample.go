package main

import "fmt"

type User struct {
	Name string
}

func updateUserName(u *User, name string) {
	u.Name = name
}

func main() {
	//v1
	// user1 := &User{Name: "Budi"}
	// updateUserName(user1, "Joko")
	
	//v2
	user1 := User{Name: "Budi"}
	fmt.Println(user1)
	updateUserName(&user1, "Joko")
	fmt.Println(user1)
}
