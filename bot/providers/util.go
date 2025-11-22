package providers

import (
	"crypto/rand"
	"io"
)

// GenerateRandomKey creates a random key with the given length in bytes.
// On failure, returns nil.
//
// Note that keys created using `GenerateRandomKey()` are not automatically
// persisted. New keys will be created when the application is restarted, and
// previously issued cookies will not be able to be decoded.
//
// Callers should explicitly check for the possibility of a nil return, treat
// it as a failure of the system random number generator, and not continue.
func GenerateRandomKey(length int) []byte {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}
