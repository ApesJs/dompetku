package auth

import (
	"fmt"
	"group-project/db_connection"
	"group-project/helper"
	"group-project/main/main_menu"
	"group-project/main/profile"
	"log"
)

func Login() {
	// KONEKSI DATABASE
	db, err := db_connection.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// LOGIN
	var username string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	password, err := helper.SensorPassword("Password: ")
	if err != nil {
		log.Fatal(err)
	}

	//clear layar console
	helper.ClearConsole()

	//query database
	query := "SELECT COUNT(*) FROM users WHERE username = ? AND password = ?"
	var count int
	err = db.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	var menu int

	if count > 0 {

		main_menu.MainMenu()
		fmt.Print("Select Menu : ")
		fmt.Scan(&menu)

		helper.ClearConsole()

		if menu == 1 {
			profile.ReadAccount(username, db)
		} else if menu == 2 {
			fmt.Println("feature is still under development")
		} else if menu == 3 {
			fmt.Println("feature is still under development")
		} else if menu == 4 {
			fmt.Println("feature is still under development")
		} else if menu == 5 {
			Login()
		}
	} else {
		fmt.Println("")
		fmt.Println("Login failed, the username or password you entered is incorrect !")
		fmt.Println("")
		fmt.Println("1. Login")
		fmt.Println("2. Forgot Password")
		fmt.Println("")
		fmt.Println("0. Exit")
		fmt.Println("")
		fmt.Print("Select Menu : ")
		fmt.Scan(&menu)

		helper.ClearConsole()

		if menu == 1 {
			Login()
		} else if menu == 2 {
			fmt.Println("feature is still under development")
		} else if menu == 3 {
			fmt.Println("feature is still under development")
		}
	}
}
