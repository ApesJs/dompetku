package profile

import (
	"database/sql"
	"fmt"
	"group-project/helper"
	"log"
)

func ReadAccount(username string, db *sql.DB) {
	query := "SELECT username, email, fullname, balance, address, phone_number, date_of_birth FROM users WHERE username = ?"
	var (
		storedUsername string
		email          string
		fullname       string
		balance        int
		address        string
		phone_number   string
		date_of_birth  string
	)
	err := db.QueryRow(query, username).Scan(&storedUsername, &email, &fullname, &balance, &address, &phone_number, &date_of_birth)
	if err != nil {
		log.Fatal(err)
	}

	formattedBalance := "Rp " + helper.FormatRupiah(balance)

	fmt.Printf("%-15s: %s\n", "Username", storedUsername)
	fmt.Printf("%-15s: %s\n", "Email", email)
	fmt.Printf("%-15s: %s\n", "Full Name", fullname)
	fmt.Printf("%-15s: %s\n", "Balance", formattedBalance)
	fmt.Printf("%-15s: %s\n", "Address", address)
	fmt.Printf("%-15s: %s\n", "Phone Number", phone_number)
	fmt.Printf("%-15s: %s\n", "Date of Birth", date_of_birth)

}
