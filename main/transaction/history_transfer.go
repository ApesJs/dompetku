package transaction

import (
	"database/sql"
	"fmt"
	"log"
)

func HistoryTransfer(username string, db *sql.DB) {
	defer db.Close()

	fmt.Println("Transfer History:")
	fmt.Println("-----------------")

	query := "SELECT sender_id, receiver_id, amount, transaction_type, message FROM transactions WHERE sender_id = ? OR receiver_id = ? ORDER BY transaction_timestamp DESC"
	rows, err := db.Query(query, username, username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var senderID, receiverID, transactionType, message string
		var amount int
		err := rows.Scan(&senderID, &receiverID, &amount, &transactionType, &message)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("Sender: %s, Receiver: %s, Amount: %d, Type: %s, Message: %s\n", senderID, receiverID, amount, transactionType, message)
	}
}
