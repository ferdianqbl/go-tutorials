package main

import (
	"fmt"
	"go-basic/package-initialization/database"
	// _ "go-basic/package-initialization/database" // to force package initialization without use function/value of the package
)

func main() {
	database.TestPackageInit()
	fmt.Println(database.GetDatabase())
}