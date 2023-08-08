package helper

import (
	"strconv"
)

func FormatRupiah(amount int) string {
	formatted := strconv.Itoa(amount)
	var result string
	for i := len(formatted) - 1; i >= 0; i-- {
		result = string(formatted[i]) + result
		if (len(formatted)-i)%3 == 0 && i > 0 {
			result = "." + result
		}
	}

	return result
}
