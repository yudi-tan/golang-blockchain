package utils

import (
	"strconv"
)

func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}
