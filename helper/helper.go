package helper

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

func ClearConsole() {
	clear := exec.Command("cmd", "/c", "cls")
	clear.Stdout = os.Stdout
	clear.Run()
}

func FormatDateOfBirth(date string, chanDate chan string) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
	}
	chanDate <- t.Format("2 January 2006")
}

func FormatDateTime(timestamp string, chanDateTime chan string) {
	t, err := time.Parse("2006-01-02 15:04:05", timestamp)
	if err != nil {
		log.Fatal(err)
	}
	chanDateTime <- t.Format("02 Jan 2006 15:04:05")
}

func FormatPhoneNumber(phoneNumber string, chanPhone chan string) {
	// Menghapus semua karakter non-digit dari nomor telepon
	phoneNumber = strings.Join(strings.FieldsFunc(phoneNumber, func(r rune) bool {
		return !('0' <= r && r <= '9')
	}), "")

	// Menambahkan kode negara dan spasi pada posisi tertentu
	formattedNumber := fmt.Sprintf("+62 %s %s %s",
		phoneNumber[1:4], phoneNumber[4:8], phoneNumber[8:])

	chanPhone <- formattedNumber
}

func FormatRupiah(amount int, chanRupiah chan string) {
	formatted := strconv.Itoa(amount)
	var result string
	for i := len(formatted) - 1; i >= 0; i-- {
		result = string(formatted[i]) + result
		if (len(formatted)-i)%3 == 0 && i > 0 {
			result = "." + result
		}
	}

	chanRupiah <- result
}

func SensorPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

func CloseChannels(channels ...interface{}) {
	for _, ch := range channels {
		switch v := ch.(type) {
		case chan string:
			close(v)
		case chan int:
			close(v)
		}
	}
}
