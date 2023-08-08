package main

import (
	"fmt"
	"group-project/db_connection"
	"group-project/helper"
	"group-project/main/auth"
	"log"
)

func main() {
	//KONEKSI DATABASE
	db, err := db_connection.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//MENU
	var menu int

	fmt.Println("")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("")
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&menu)

	helper.ClearConsole()

	if menu == 1 {
		auth.Login()
	} else {
		//REGISTER
	}
}
