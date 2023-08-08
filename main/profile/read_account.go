package profile

import (
	"database/sql"
	"fmt"
	"group-project/helper"
	"group-project/main/main_menu"
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

	var menu int
	fmt.Println("")
	fmt.Println("1 Edit My Profile")
	fmt.Println("2 Delete My Profile")
	fmt.Println("3 Back")
	fmt.Println("")
	fmt.Println("0 Exit")
	fmt.Println("")
	fmt.Print("Select Menu : ")
	fmt.Scan(&menu)

	helper.ClearConsole()

	if menu == 1 {
		fmt.Println("EDIT COMMING SOON")
	} else if menu == 2 {
		fmt.Println("DELETE COMMING SOON")
	} else if menu == 3 {
		main_menu.MainMenu()
		fmt.Print("Select Menu : ")
		fmt.Scan(&menu)

		if menu == 1 {
			ReadAccount(username, db)
		}
	}

}
