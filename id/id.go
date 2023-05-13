package id

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"time"
)

// Token - generates random secure string
func Token() (string, error) {
	var (
		buf [8]byte
		b   = make([]byte, 16)
	)
	binary.BigEndian.PutUint64(buf[:], uint64(time.Now().UnixNano()))
	_, err := rand.Read(b)
	if err == nil {
		b = append(b, buf[:]...)
		return base64.StdEncoding.EncodeToString(b), nil
	}
	return "", err
}
