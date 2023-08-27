package profile

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/ApesJs/dompetku/db_connection"
	"github.com/ApesJs/dompetku/helper"
	"github.com/ApesJs/dompetku/main/main_menu"
	"github.com/ApesJs/dompetku/main/search_user"
	"github.com/ApesJs/dompetku/main/transaction"

	"log"
	"time"
)

type Users struct {
	RWMutex                              sync.RWMutex
	Username, Password, Email, Fullname  string
	Balance                              int
	Address, Phone_number, Date_of_birth string
}

func (user *Users) GetUser(username string, db *sql.DB, chanUsers chan string) {
	user.RWMutex.Lock()
	query := "SELECT username, email, fullname, balance, address, phone_number, date_of_birth FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&user.Username, &user.Email, &user.Fullname, &user.Balance, &user.Address, &user.Phone_number, &user.Date_of_birth)
	if err != nil {
		log.Fatal(err)
	}

	go helper.FormatRupiah(user.Balance, chanUsers)
	formattedBalance := <-chanUsers

	go helper.FormatPhoneNumber(user.Phone_number, chanUsers)
	formatedPhoneNumber := <-chanUsers

	go helper.FormatDateOfBirth(user.Date_of_birth, chanUsers)
	formattedDateOfBirth := <-chanUsers

	fmt.Printf("%-15s: %s\n", "Username", user.Username)
	fmt.Printf("%-15s: %s\n", "Email", user.Email)
	fmt.Printf("%-15s: %s\n", "Full Name", user.Fullname)
	fmt.Printf("%-15s: %s\n", "Balance", "Rp "+formattedBalance)
	fmt.Printf("%-15s: %s\n", "Address", user.Address)
	fmt.Printf("%-15s: %s\n", "Phone Number", formatedPhoneNumber)
	fmt.Printf("%-15s: %s\n", "Date of Birth", formattedDateOfBirth)
	user.RWMutex.Unlock()
}

func ReadAccount(username string, db *sql.DB) {
	defer db_connection.PutDB(db)
	UserData := Users{}
	channelMainMenu := make(chan string)
	chanUsers := make(chan string)
	defer close(chanUsers)

	go UserData.GetUser(username, db, chanUsers)

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
		go main_menu.MainMenu(channelMainMenu)

		for mainmenu := range channelMainMenu {
			fmt.Println(mainmenu)
		}
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
