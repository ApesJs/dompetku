package search_user

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID             string
	Username       string
	Password       string
	Email          string
	Fullname       string
	Balance        int
	ProfilePicture string
	Address        string
	PhoneNumber    string
	DateOfBirth    string
	status_delete  int
}

func SearchUser(db *sql.DB) {

	fmt.Println("All Full Names:")
	rows, err := db.Query("SELECT fullname FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var fullname string
		err := rows.Scan(&fullname)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(fullname)
	}

	var searchFullname string
	fmt.Print("Enter fullname: ")
	fmt.Scanln(&searchFullname)

	query := "SELECT id, username, password, email, fullname, balance, profile_picture, address, phone_number, date_of_birth FROM users WHERE fullname LIKE ? AND status_delete is null"
	rows, err = db.Query(query, "%"+searchFullname+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Search Results:")

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID, &user.Username, &user.Password, &user.Email, &user.Fullname,
			&user.Balance, &user.ProfilePicture, &user.Address, &user.PhoneNumber, &user.DateOfBirth,
		)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("User ID: %s\n", user.ID)
		fmt.Printf("Username: %s\n", user.Username)
		fmt.Printf("Password: %s\n", user.Password)
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Printf("Fullname: %s\n", user.Fullname)
		fmt.Printf("Balance: %d\n", user.Balance)
		fmt.Printf("Profile Picture: %s\n", user.ProfilePicture)
		fmt.Printf("Address: %s\n", user.Address)
		fmt.Printf("Phone Number: %s\n", user.PhoneNumber)
		fmt.Printf("Date of Birth: %s\n", user.DateOfBirth)
	}
}
