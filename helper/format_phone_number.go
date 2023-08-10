package helper

import (
	"fmt"
	"strings"
)

func FormatPhoneNumber(phoneNumber string) string {
	// Menghapus semua karakter non-digit dari nomor telepon
	phoneNumber = strings.Join(strings.FieldsFunc(phoneNumber, func(r rune) bool {
		return !('0' <= r && r <= '9')
	}), "")

	// Menambahkan kode negara dan spasi pada posisi tertentu
	formattedNumber := fmt.Sprintf("+62 %s %s %s",
		phoneNumber[1:4], phoneNumber[4:8], phoneNumber[8:])

	return formattedNumber
}
