package profile

import (
	"bufio"
	"database/sql"
	"fmt"
	"group-project/helper"
	"log"
	"os"
	"strings"
	"time"
)

func ValidateDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

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

					fmt.Print("")
					fmt.Print("Do you want to exit? (Yes/No) : ")
					fmt.Print("")
					fmt.Scan(&answer)

					helper.ClearConsole()

					if answer == "No" || answer == "no" {
						ReadAccount(username, db)
					} else {
						fmt.Print("Exit....")
						time.Sleep(2 * time.Second)
						helper.ClearConsole()
					}
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
					time.Sleep(2 * time.Second)
					helper.ClearConsole()
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
					time.Sleep(2 * time.Second)
					helper.ClearConsole()
					break
				}
			}
		} else if menu == 4 {
			for {
				fmt.Println("")
				fmt.Println("Remember !, change your name carefully !")
				fmt.Println("")
				fmt.Print("Enter your new name : ")
				fmt.Scan(&UpdateUser.Fullname)

				takeMyFullname := "SELECT fullname FROM users WHERE username = ?"
				var myFullname string
				errTakeMyFullname := db.QueryRow(takeMyFullname, username).Scan(&myFullname)
				if errTakeMyFullname != nil {
					log.Fatal(errTakeMyFullname)
				}

				if UpdateUser.Fullname == myFullname {
					fmt.Println("")
					fmt.Println("The name you entered is the same !")
					fmt.Print("Are you sure you want to keep changing your name? (Yes/No) : ")
					fmt.Scan(&answer)

					if answer == "No" || answer == "no" {
						break
					}

					helper.ClearConsole()
				} else {
					//query database
					query := "UPDATE users SET fullname = ? WHERE username = ?"
					_, err := db.Exec(query, UpdateUser.Fullname, username)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("")
					fmt.Println("Your name has been successfully updated !")
					fmt.Println("")
					time.Sleep(2 * time.Second)
					helper.ClearConsole()
					break
				}
			}
		} else if menu == 5 {
			fmt.Println("")
			fmt.Println("Remember !, change your address carefully !")
			fmt.Println("")
			fmt.Print("Enter your new address : ")

			// Membaca input teks panjang hingga ditemukan baris kosong (hanya '.')
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				line := scanner.Text()
				if line == "." {
					break
				}
				UpdateUser.Address += line + "\n"
			}

			// Menghapus karakter terakhir (newline) dari address
			UpdateUser.Address = strings.TrimSuffix(UpdateUser.Address, "\n")

			//query database
			query := "UPDATE users SET address = ? WHERE username = ?"
			_, err := db.Exec(query, UpdateUser.Address, username)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("")
			fmt.Println("Your address has been successfully updated !")
			fmt.Println("")
			time.Sleep(2 * time.Second)
			helper.ClearConsole()
		} else if menu == 6 {
			for {
				fmt.Println("")
				fmt.Println("Remember !, change your phone number carefully !")
				fmt.Println("")
				fmt.Print("Enter your new phone number : ")
				fmt.Scan(&UpdateUser.Phone_number)

				checkPhoneNumber := "SELECT COUNT(*) FROM users WHERE phone_number = ?"
				var count int
				errCheckPhoneNumber := db.QueryRow(checkPhoneNumber, UpdateUser.Phone_number).Scan(&count)
				if errCheckPhoneNumber != nil {
					log.Fatal(errCheckPhoneNumber)
				}

				takeMyPhoneNumber := "SELECT phone_number FROM users WHERE username = ?"
				var myPhoneNumber string
				errTakeMyPhoneNumber := db.QueryRow(takeMyPhoneNumber, username).Scan(&myPhoneNumber)
				if errTakeMyPhoneNumber != nil {
					log.Fatal(errTakeMyPhoneNumber)
				}

				if UpdateUser.Phone_number == myPhoneNumber {
					fmt.Println("")
					fmt.Println("The phone number you entered is the same !")
					fmt.Print("Are you sure you want to keep changing your phone number? (Yes/No) : ")
					fmt.Scan(&answer)

					if answer == "No" || answer == "no" {
						break
					}

					helper.ClearConsole()
				} else if count > 0 {
					fmt.Println("")
					fmt.Println("The phone number you entered is already in use, try again !")
					time.Sleep(2 * time.Second)
					helper.ClearConsole()
				} else if count == 0 {
					//query database
					query := "UPDATE users SET phone_number = ? WHERE username = ?"
					_, err := db.Exec(query, UpdateUser.Phone_number, username)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("")
					fmt.Println("Your phone number has been successfully updated !")
					fmt.Println("")
					time.Sleep(2 * time.Second)
					helper.ClearConsole()
					break
				}
			}
		} else if menu == 7 {
			for {
				fmt.Println("")
				fmt.Println("Remember !, The format for filling in the date of birth is Year-Month-Date !")
				fmt.Println("")
				fmt.Print("Enter your new date of birth : ")
				fmt.Scan(&UpdateUser.Date_of_birth)

				takeMyDateOfBirth := "SELECT date_of_birth FROM users WHERE username = ?"
				var myDateOfBirth string
				errTakeMyDateOfBirth := db.QueryRow(takeMyDateOfBirth, username).Scan(&myDateOfBirth)
				if errTakeMyDateOfBirth != nil {
					log.Fatal(errTakeMyDateOfBirth)
				}

				if UpdateUser.Date_of_birth == myDateOfBirth {
					fmt.Println("")
					fmt.Println("The date of birth you entered is the same !")
					fmt.Print("Are you sure you want to keep changing your date of birth? (Yes/No) : ")
					fmt.Scan(&answer)

					if answer == "No" || answer == "no" {
						break
					}

					helper.ClearConsole()
				} else {
					//query database
					query := "UPDATE users SET date_of_birth = ? WHERE username = ?"
					_, err := db.Exec(query, UpdateUser.Date_of_birth, username)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("")
					fmt.Println("Your date of birth has been successfully updated !")
					fmt.Println("")
					time.Sleep(2 * time.Second)
					helper.ClearConsole()
					break
				}
			}
		} else if menu == 8 {
			fmt.Println("")
			fmt.Println("Remember !, change your profile carefully !")
			fmt.Println("")
			fmt.Printf("%-15s: ", "Username")
			fmt.Scan(&UpdateUser.Username)
			fmt.Printf("%-15s: ", "Password")
			fmt.Scan(&UpdateUser.Password)
			fmt.Printf("%-15s: ", "Email")
			fmt.Scan(&UpdateUser.Email)
			fmt.Printf("%-15s: ", "Full Name")
			scanFullname := bufio.NewScanner(os.Stdin)
			for scanFullname.Scan() {
				line := scanFullname.Text()
				if line == "." {
					break
				}
				UpdateUser.Fullname += line
			}
			UpdateUser.Fullname = strings.TrimSuffix(UpdateUser.Fullname, "\n")
			fmt.Printf("%-15s: ", "Address")
			scanPassword := bufio.NewScanner(os.Stdin)
			for scanPassword.Scan() {
				line := scanPassword.Text()
				if line == "." {
					break
				}
				UpdateUser.Address += line
			}
			UpdateUser.Address = strings.TrimSuffix(UpdateUser.Address, "\n")
			fmt.Printf("%-15s: ", "Phone Number")
			fmt.Scan(&UpdateUser.Phone_number)
			fmt.Printf("%-15s: ", "Date of Birth")
			fmt.Scan(&UpdateUser.Date_of_birth)

			query := "UPDATE users SET username = ?, password = ?, email = ?, fullname = ?, address = ?, phone_number = ?, date_of_birth = ? WHERE username = ?"
			_, err := db.Exec(query, UpdateUser.Username, UpdateUser.Password, UpdateUser.Email, UpdateUser.Fullname, UpdateUser.Address, UpdateUser.Phone_number, UpdateUser.Date_of_birth, username)
			if err != nil {
				log.Fatal(err)
			} else if err == nil {
				fmt.Println("")
				fmt.Println("Your profile has been successfully updated !")
				fmt.Println("")
				time.Sleep(2 * time.Second)
				helper.ClearConsole()
				ReadAccount(username, db)
			}

		} else if menu == 0 {
			ReadAccount(username, db)
		}
	}

}
