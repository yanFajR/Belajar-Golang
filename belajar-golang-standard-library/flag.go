package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")

	flag.Parse()

	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
	fmt.Println("Host :", *host)
}