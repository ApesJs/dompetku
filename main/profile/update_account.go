package profile

import (
	"database/sql"
	"fmt"
	"group-project/helper"
	"log"
	"time"
)

func UpdateAccount(username string, db *sql.DB) {
	UpdateUser := Users{}

	fmt.Println("")
	fmt.Println("Edit Your Profile")
	fmt.Println("1. Username")
	fmt.Println("2. Password")
	fmt.Println("3. Email")
	fmt.Println("4. Name")
	fmt.Println("5. Address")
	fmt.Println("6. Phone Number")
	fmt.Println("7. Date of Birth")
	fmt.Println("")
	fmt.Println("8. Edit All My Profile")
	fmt.Println("")
	fmt.Println("0. Back")
	fmt.Println("")

	var menu int
	fmt.Print("Select Menu : ")
	fmt.Scan(&menu)

	helper.ClearConsole()

	var answer string
	fmt.Print("Are you sure (Yes/No) ? ")
	fmt.Scan(&answer)

	helper.ClearConsole()

	if answer == "Yes" || answer == "yes" {
		if menu == 1 {
			for {
				fmt.Println("")
				fmt.Println("Remember !, username can only be changed once a month !")
				fmt.Println("")
				fmt.Print("Enter your new username : ")
				fmt.Scan(&UpdateUser.Username)

				checkUsername := "SELECT COUNT(*) FROM users WHERE username = ?"
				var count int
				errCheckUsername := db.QueryRow(checkUsername, UpdateUser.Username).Scan(&count)
				if errCheckUsername != nil {
					log.Fatal(errCheckUsername)
				}

				if UpdateUser.Username == username {
					fmt.Println("")
					fmt.Println("The username you entered is the same !")
					fmt.Print("Are you sure you want to keep changing your username? (Yes/No) : ")
					fmt.Scan(&answer)

					if answer == "No" || answer == "no" {
						helper.ClearConsole()
						ReadAccount(username, db)
					}

					helper.ClearConsole()
				} else if count > 0 {
					fmt.Println("")
					fmt.Println("The username you entered is already in use, try again !")
					time.Sleep(2 * time.Second)
					helper.ClearConsole()
				} else if count == 0 {
					//query database
					query := "UPDATE users SET username = ? WHERE username = ?"
					_, err := db.Exec(query, UpdateUser.Username, username)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("")
					fmt.Println("Your username has been successfully updated !")
					fmt.Println("")
					time.Sleep(2 * time.Second)
					helper.ClearConsole()
					break
				}
			}
		} else if menu == 2 {
			var tryPass string
			for {
				fmt.Println("")
				fmt.Println("Remember !, Be careful when setting a new password !")
				fmt.Println("")
				fmt.Print("Enter your new password : ")
				fmt.Scan(&UpdateUser.Password)
				fmt.Print("re-enter password : ")
				fmt.Scan(&tryPass)

				if tryPass != UpdateUser.Password {
					fmt.Println("")
					fmt.Println("The passwords entered are not the same, try again !")
					fmt.Println("")
					time.Sleep(2 * time.Second)
					helper.ClearConsole()
				} else {
					//query database
					query := "UPDATE users SET password = ? WHERE username = ?"
					_, err := db.Exec(query, UpdateUser.Password, username)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("")
					fmt.Println("Your password has been successfully updated !")
					fmt.Println("")

					break
				}
			}
		} else if menu == 3 {
			for {
				fmt.Println("")
				fmt.Println("Remember !, change your email carefully !")
				fmt.Println("")
				fmt.Print("Enter your new email : ")
				fmt.Scan(&UpdateUser.Email)

				checkEmail := "SELECT COUNT(*) FROM users WHERE email = ?"
				var count int
				errCheckEmail := db.QueryRow(checkEmail, UpdateUser.Email).Scan(&count)
				if errCheckEmail != nil {
					log.Fatal(errCheckEmail)
				}

				takeMyEmail := "SELECT email FROM users WHERE username = ?"
				var myEmail string
				errTakeMyEmail := db.QueryRow(takeMyEmail, username).Scan(&myEmail)
				if errTakeMyEmail != nil {
					log.Fatal(errTakeMyEmail)
				}

				if UpdateUser.Email == myEmail {
					fmt.Println("")
					fmt.Println("The email you entered is the same !")
					fmt.Print("Are you sure you want to keep changing your email? (Yes/No) : ")
					fmt.Scan(&answer)

					if answer == "No" || answer == "no" {
						break
					}

					helper.ClearConsole()
				} else if count > 0 {
					fmt.Println("")
					fmt.Println("The email you entered is already in use, try again !")
					time.Sleep(2 * time.Second)
					helper.ClearConsole()
				} else if count == 0 {
					//query database
					query := "UPDATE users SET email = ? WHERE username = ?"
					_, err := db.Exec(query, UpdateUser.Email, username)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("")
					fmt.Println("Your email has been successfully updated !")
					fmt.Println("")

					break
				}
			}
		} else if menu == 4 {
			fmt.Println("")
		} else if menu == 5 {
			fmt.Println("")
		} else if menu == 6 {
			fmt.Println("")
		} else if menu == 7 {
			fmt.Println("")
		} else if menu == 8 {
			fmt.Println("")
		} else if menu == 0 {
			ReadAccount(username, db)
		}
	} else {
		UpdateAccount(username, db)
	}

}
