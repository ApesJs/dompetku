package helper

import (
	"log"
	"time"
)

func FormatDateOfBirth(date string) string {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
	}
	return t.Format("2 January 2006")
}
