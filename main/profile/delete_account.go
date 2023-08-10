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
	// db, err := sql.Open("mysql", "root:1Muhammad@tcp(127.0.0.1:3306)/db_project1")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db, err := db_connection.ConnectDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	targetUser := User{}
	fmt.Print("masukkan password:")
	fmt.Scanln(&targetUser.Password)
	//var success int
	query := "UPDATE users SET status_delete = 1 WHERE Password = ?"
	_, err := db.Exec(query, targetUser.Password)
	if err != nil {
		log.Fatal(err)

		// }else if err == nil {
		// 	success=1
		// }
		// if success == 1 {
		// 	auth.Login()
	}
	fmt.Println("delete berhasil")
}
