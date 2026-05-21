package database

import "fmt"

var connection string

// this function auto run when package initialization / the user use this package/all function once time
func init() {
	fmt.Println("Package initialization")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}

func TestPackageInit(){
	fmt.Println("Test Package Initialization")
}