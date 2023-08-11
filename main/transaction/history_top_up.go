package transaction

import (
	"database/sql"
	"fmt"
	"group-project/helper"
	"log"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type TransactionHistory struct {
	SenderID        string
	Amount          int
	TransactionType string
	Message         string
	TransactionTime string
}

func FormatDateTime(timestamp string) string {
	t, err := time.Parse("2006-01-02 15:04:05", timestamp)
	if err != nil {
		log.Fatal(err)
	}
	return t.Format("02 Jan 2006 15:04:05")
}

func HistoryTopUp(username string, db *sql.DB) {
	var history string
	transactionList := []TransactionHistory{}

	query := "SELECT sender_id, amount, transaction_type, message, transaction_timestamp FROM transactions WHERE sender_id = ?"
	rows, err := db.Query(query, username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction TransactionHistory
		err := rows.Scan(&transaction.SenderID, &transaction.Amount, &transaction.TransactionType, &transaction.Message, &history)
		if err != nil {
			log.Fatal(err)
		}
		transaction.TransactionTime = FormatDateTime(history)
		transactionList = append(transactionList, transaction)
	}

	fmt.Println("")
	fmt.Println("This is your top up history")
	fmt.Println("")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Sender ID", "Amount", "Transaction Type", "Message", "Transaction Time"})

	for _, transaction := range transactionList {
		table.Append([]string{transaction.SenderID, fmt.Sprintf("Rp " + helper.FormatRupiah(transaction.Amount)), transaction.TransactionType, transaction.Message, transaction.TransactionTime})
	}

	table.Render()
}
