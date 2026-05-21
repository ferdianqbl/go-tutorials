package main

import "fmt"



type User struct{
	Name string
}


func main() {

	user1:= new(User) // pointer to empty struct
	user2:= user1

	user2.Name = "Budi"


	fmt.Println(user1)
	fmt.Println(user2)

}