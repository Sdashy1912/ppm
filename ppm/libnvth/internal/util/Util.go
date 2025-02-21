package util

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 sha256
func SHA256(b []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(b))
}
