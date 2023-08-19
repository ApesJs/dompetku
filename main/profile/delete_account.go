package profile

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID             int
	Username       string
	Password       string
	Email          string
	Fullname       string
	Balance        int
	ProfilePicture string
	Address        string
	PhoneNumber    string
	DateOfBirth    string
}

func DeleteAccount(password string, db *sql.DB) {
	defer db.Close()
	targetUser := User{}
	fmt.Print("masukkan username:")
	fmt.Scanln(&targetUser.Username)
	fmt.Print("masukkan password:")
	fmt.Scanln(&targetUser.Password)

	query := "UPDATE users SET status_delete = 1 WHERE username =? AND password = ?"
	_, err := db.Exec(query, targetUser.Username, targetUser.Password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete berhasil")
}
