package utils

import (
	"crypto/sha256"
	"fmt"
)

func Sha256WithSalt(s string, salt string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s+salt)))
}
