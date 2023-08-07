package main

import (
	"fmt"
	"group-project/db_connection"
	"log"
)

func main() {
	//KONEKSI DATABASE
	db, err := db_connection.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//LOGIN
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	query := "SELECT COUNT(*) FROM users WHERE username = ? AND password = ?"
	var count int
	err = db.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		fmt.Println("Login berhasil!")
	} else {
		fmt.Println("Login gagal.")
	}
}
