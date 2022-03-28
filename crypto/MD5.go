package crypto

import (
	"crypto/md5"
	"fmt"
)

func CryptMD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}
