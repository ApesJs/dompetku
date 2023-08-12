package profile

import (
	"database/sql"
	"fmt"
	"group-project/helper"
	"group-project/main/search_user"
	"group-project/main/transaction"

	"log"
	"time"
)

type Users struct {
	Username, Password, Email, Fullname  string
	Balance                              int
	Address, Phone_number, Date_of_birth string
}

func ReadAccount(username string, db *sql.DB) {
	query := "SELECT username, email, fullname, balance, address, phone_number, date_of_birth FROM users WHERE username = ?"
	UserData := Users{}
	err := db.QueryRow(query, username).Scan(&UserData.Username, &UserData.Email, &UserData.Fullname, &UserData.Balance, &UserData.Address, &UserData.Phone_number, &UserData.Date_of_birth)
	if err != nil {
		log.Fatal(err)
	}

	formattedBalance := "Rp " + helper.FormatRupiah(UserData.Balance)
	formatedPhoneNumber := helper.FormatPhoneNumber(UserData.Phone_number)
	formattedDateOfBirth := helper.FormatDateOfBirth(UserData.Date_of_birth)

	fmt.Printf("%-15s: %s\n", "Username", UserData.Username)
	fmt.Printf("%-15s: %s\n", "Email", UserData.Email)
	fmt.Printf("%-15s: %s\n", "Full Name", UserData.Fullname)
	fmt.Printf("%-15s: %s\n", "Balance", formattedBalance)
	fmt.Printf("%-15s: %s\n", "Address", UserData.Address)
	fmt.Printf("%-15s: %s\n", "Phone Number", formatedPhoneNumber)
	fmt.Printf("%-15s: %s\n", "Date of Birth", formattedDateOfBirth)

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
		UpdateAccount(username, db)
	} else if menu == 2 {
		DeleteAccount(username, db)
	} else if menu == 3 {
		// main_menu.MainMenu()
		fmt.Print("Select Menu : ")
		fmt.Scan(&menu)

		helper.ClearConsole()

		if menu == 1 {
			ReadAccount(username, db)
		} else if menu == 2 {
			search_user.SearchUser(db)
		} else if menu == 3 {
			fmt.Println("")
			fmt.Println("1 Top Up")
			fmt.Println("2 Transfer")
			fmt.Println("")

			fmt.Print("Select Menu : ")
			fmt.Scan(&menu)

			helper.ClearConsole()

			if menu == 1 {
				transaction.TopUp(username, db)
			} else if menu == 2 {
				fmt.Print("Transfer")
			}
		} else if menu == 4 {
			fmt.Println("")
			fmt.Println("1 History Top Up")
			fmt.Println("2 History Transfer")
			fmt.Println("")

			fmt.Print("Select Menu : ")
			fmt.Scan(&menu)

			helper.ClearConsole()

			if menu == 1 {
				transaction.HistoryTopUp(username, db)
			} else if menu == 2 {
				fmt.Print("History Transfer")
			}
		} else if menu == 5 {
			fmt.Println("auth.Login()")
		} else if menu == 0 {
			fmt.Println("Exit....")
			time.Sleep(2 * time.Second)
			helper.ClearConsole()
		}
	} else if menu == 0 {
		fmt.Println("Exit....")
		time.Sleep(2 * time.Second)
		helper.ClearConsole()
	}

}
