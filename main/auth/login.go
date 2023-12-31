package auth

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/ApesJs/dompetku/db_connection"
	"github.com/ApesJs/dompetku/helper"
	"github.com/ApesJs/dompetku/main/main_menu"
	"github.com/ApesJs/dompetku/main/profile"
	"github.com/ApesJs/dompetku/main/search_user"
	"github.com/ApesJs/dompetku/main/transaction"

	"log"
	"time"
)

func Login() {
	channelMainMenu := make(chan string)

	// KONEKSI DATABASE
	var (
		db *sql.DB
		wg sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		db = db_connection.GetDB()
		defer db_connection.PutDB(db)
	}()
	wg.Wait()

	// LOGIN
	var username string
	fmt.Println("")
	fmt.Println("Enter your username and password")
	fmt.Println("")
	fmt.Print("Username: ")
	fmt.Scan(&username)
	password, err := helper.SensorPassword("Password: ")
	if err != nil {
		log.Fatal(err)
	}

	//clear layar console
	helper.ClearConsole()

	//query database
	query := "SELECT COUNT(*) FROM users WHERE username = ? AND password = ? AND status_delete = 0"
	var count int
	err = db.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	var menu int

	if count > 0 {
		go main_menu.MainMenu(channelMainMenu)

		for mainmenu := range channelMainMenu {
			fmt.Println(mainmenu)
		}

		fmt.Print("Select Menu : ")
		fmt.Scan(&menu)

		helper.ClearConsole()

		if menu == 1 {
			profile.ReadAccount(username, db)
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
				transaction.Transfer(username, db)

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
				transaction.HistoryTransfer(username, db)

			}
		} else if menu == 5 {
			Login()
		} else if menu == 0 {
			fmt.Println("Exit....")
			time.Sleep(2 * time.Second)
			helper.ClearConsole()
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
		} else if menu == 0 {
			fmt.Println("Exit....")
			time.Sleep(2 * time.Second)
			helper.ClearConsole()
		}
	}
}
