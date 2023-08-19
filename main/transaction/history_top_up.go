package transaction

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ApesJs/dompetku/helper"

	"github.com/olekukonko/tablewriter"
)

type TransactionHistory struct {
	SenderID        string
	Amount          int
	TransactionType string
	Message         string
	TransactionTime string
}

func HistoryTopUp(username string, db *sql.DB) {
	defer db.Close()
	chanRupiah := make(chan string)
	chanDateTime := make(chan string)
	defer helper.CloseChannels(chanRupiah, chanDateTime)

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
		go helper.FormatRupiah(transaction.Amount, chanRupiah)
		go helper.FormatDateTime(history, chanDateTime)
		transactionList = append(transactionList, transaction)
	}

	fmt.Println("")
	fmt.Println("This is your top up history")
	fmt.Println("")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Sender ID", "Amount", "Transaction Type", "Message", "Transaction Time"})

	for _, transaction := range transactionList {
		table.Append([]string{transaction.SenderID, fmt.Sprintf("Rp " + <-chanRupiah), transaction.TransactionType, transaction.Message, <-chanDateTime})
	}

	table.Render()
}
