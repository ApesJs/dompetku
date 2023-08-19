package transaction

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ApesJs/dompetku/helper"
)

type Transaction struct {
	RWMutex                   sync.RWMutex
	Sender_id                 string
	Amount                    int
	Transaction_type, Message string
}

func (transaction *Transaction) AddBalance(username string, db *sql.DB) {
	transaction.RWMutex.Lock()
	takeMyCurrentBalance := "SELECT balance FROM users WHERE username = ?"
	var myCurrentBalance int
	errTakeMyCurrentBalance := db.QueryRow(takeMyCurrentBalance, username).Scan(&myCurrentBalance)
	if errTakeMyCurrentBalance != nil {
		log.Println(errTakeMyCurrentBalance)
	}

	sumMyBalance := transaction.Amount + myCurrentBalance

	query := "UPDATE users SET balance = ? WHERE username = ?"
	_, err := db.Exec(query, sumMyBalance, username)
	if err != nil {
		log.Println(err)
	}
	transaction.RWMutex.Unlock()
}

func (transaction *Transaction) AddTransaction(db *sql.DB) {
	transaction.RWMutex.Lock()
	transactionInsert := "INSERT INTO transactions (sender_id, receiver_id, amount, transaction_type, message) VALUES(?,?,?,?,?)"
	_, err2 := db.Exec(transactionInsert, transaction.Sender_id, transaction.Sender_id, transaction.Amount, "Top Up", transaction.Message)
	if err2 != nil {
		log.Println(err2)
	}
	transaction.RWMutex.Unlock()
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
			go TransactionUser.AddBalance(username, db)
			go TransactionUser.AddTransaction(db)

			fmt.Println("")
			fmt.Println("Your top up was successful !")
			fmt.Println("")
			time.Sleep(2 * time.Second)
			helper.ClearConsole()
			break
		}
	}
}
