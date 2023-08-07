// auth/login.go
package auth

import (
	"fmt"
	"group-project/db_connection"
	"log"
	"os"
	"os/exec"
)

func ClearConsole() {
	clear := exec.Command("cmd", "/c", "cls")
	clear.Stdout = os.Stdout
	clear.Run()
}

func Login() {
	// KONEKSI DATABASE
	db, err := db_connection.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// LOGIN
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	ClearConsole()

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
