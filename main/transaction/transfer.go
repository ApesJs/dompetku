package transaction

import (
	"database/sql"
	"fmt"
	"group-project/helper"
	"log"
	"time"
)

type Transactions struct {
	SenderID        string
	ReceiverID      string
	Amount          int
	TransactionType string
	Message         string
}

func Transfer(username string, db *sql.DB) {
	for {
		TransactionUser := Transactions{}

		fmt.Println("")
		fmt.Print("Enter your username: ")
		fmt.Scan(&TransactionUser.SenderID)
		fmt.Print("Enter the recipient's username: ")
		fmt.Scan(&TransactionUser.ReceiverID)
		fmt.Print("Enter the transfer amount: ")
		fmt.Scan(&TransactionUser.Amount)
		fmt.Print("Enter a transaction message: ")
		fmt.Scan(&TransactionUser.Message)

		if TransactionUser.SenderID != username {
			fmt.Println("")
			fmt.Println("The username you entered is incorrect!")
			time.Sleep(2 * time.Second)
			helper.ClearConsole()
		} else {
			takeMyCurrentBalance := "SELECT balance FROM users WHERE username = ?"
			var myCurrentBalance int
			errTakeMyCurrentBalance := db.QueryRow(takeMyCurrentBalance, username).Scan(&myCurrentBalance)
			if errTakeMyCurrentBalance != nil {
				log.Fatal(errTakeMyCurrentBalance)
			}

			if myCurrentBalance < TransactionUser.Amount {
				fmt.Println("")
				fmt.Println("Insufficient balance for transfer!")
				time.Sleep(2 * time.Second)
				helper.ClearConsole()
				break
			}

			sumMyBalance := myCurrentBalance - TransactionUser.Amount

			updateSenderBalance := "UPDATE users SET balance = ? WHERE username = ?"
			_, err := db.Exec(updateSenderBalance, sumMyBalance, username)
			if err != nil {
				log.Fatal(err)
			}

			updateReceiverBalance := "UPDATE users SET balance = balance + ? WHERE username = ?"
			_, err = db.Exec(updateReceiverBalance, TransactionUser.Amount, TransactionUser.ReceiverID)
			if err != nil {
				log.Fatal(err)
			}

			transactionInsert := "INSERT INTO transactions (sender_id, receiver_id, amount, transaction_type, message) VALUES(?,?,?,?,?)"
			_, err = db.Exec(transactionInsert, TransactionUser.SenderID, TransactionUser.ReceiverID, TransactionUser.Amount, "Transfer", TransactionUser.Message)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("")
			fmt.Println("Your transfer was successful!")
			fmt.Println("")
			time.Sleep(2 * time.Second)
			helper.ClearConsole()
			break
		}
	}
}
