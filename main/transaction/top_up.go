package transaction

import (
	"database/sql"
	"fmt"
	"group-project/helper"
	"log"
	"time"
)

type Transaction struct {
	Sender_id                 string
	Amount                    int
	Transaction_type, Message string
}

func TopUp(username string, db *sql.DB) {
	for {
		TransactionUser := Transaction{}

		fmt.Println("")
		fmt.Print("Enter your username :")
		fmt.Scan(&TransactionUser.Sender_id)
		fmt.Print("Enter the top up amount :")
		fmt.Scan(&TransactionUser.Amount)
		fmt.Print("Enter a transaction message :")
		fmt.Scan(&TransactionUser.Message)

		if TransactionUser.Sender_id != username {
			fmt.Println("")
			fmt.Println("The username you entered is incorrect !")
			time.Sleep(2 * time.Second)

			helper.ClearConsole()

		} else if TransactionUser.Sender_id == username {
			takeMyCurrentBalance := "SELECT balance FROM users WHERE username = ?"
			var myCurrentBalance int
			errTakeMyCurrentBalance := db.QueryRow(takeMyCurrentBalance, username).Scan(&myCurrentBalance)
			if errTakeMyCurrentBalance != nil {
				log.Fatal(errTakeMyCurrentBalance)
			}

			sumMyBalance := TransactionUser.Amount + myCurrentBalance

			query := "UPDATE users SET balance = ? WHERE username = ?"
			_, err := db.Exec(query, sumMyBalance, username)
			if err != nil {
				log.Fatal(err)
			}

			transactionInsert := "INSERT INTO transactions (sender_id, receiver_id, amount, transaction_type, message) VALUES(?,?,?,?,?)"
			_, err2 := db.Exec(transactionInsert, TransactionUser.Sender_id, TransactionUser.Sender_id, TransactionUser.Amount, "Top Up", TransactionUser.Message)
			if err2 != nil {
				log.Fatal(err2)
			}

			fmt.Println("")
			fmt.Println("Your top up was successful !")
			fmt.Println("")
			time.Sleep(2 * time.Second)
			helper.ClearConsole()
			break
		}
	}
}
