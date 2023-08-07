package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/account_service")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
