package main

import (
	"fmt"
	"go-basic/package-initialization/database"
)

func main() {
	fmt.Println(database.GetDatabase())
	database.TestPackageInit()
}