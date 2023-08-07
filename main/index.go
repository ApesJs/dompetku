package main

import (
	"fmt"
	"group-project/db_connection"
	"log"
	"os"
	"os/exec"
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
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&menu)

	//clear layar console
	clear := exec.Command("cmd", "/c", "cls")
	clear.Stdout = os.Stdout
	clear.Run()

	if menu == 1 {
		//LOGIN
		var username, password string
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		//clear layar console
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()

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
	} else {
		//REGISTER
	}
}
