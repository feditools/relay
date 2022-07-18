package util

import (
	"crypto/aes"
	gocipher "crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
	"sync"

	"github.com/feditools/relay/internal/config"
	"github.com/spf13/viper"
)

var (
	cryptoKey     [32]byte
	cryptoKeyOnce sync.Once

	ErrDataTooSmall = errors.New("data too small")
)

func Decrypt(b []byte) ([]byte, error) {
	gcm, err := getCrypto()
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(b) < nonceSize {
		return nil, ErrDataTooSmall
	}

	nonce, ciphertext := b[:nonceSize], b[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func Encrypt(b []byte) ([]byte, error) {
	gcm, err := getCrypto()
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, b, nil), nil
}

func getCrypto() (gocipher.AEAD, error) {
	cipher, err := aes.NewCipher(getKey())
	if err != nil {
		return nil, err
	}

	gcm, err := gocipher.NewGCM(cipher)
	if err != nil {
		return nil, err
	}

	return gcm, nil
}

func getKey() []byte {
	cryptoKeyOnce.Do(func() {
		cryptoKey = sha256.Sum256([]byte(viper.GetString(config.Keys.EncryptionKey)))
	})

	return cryptoKey[:]
}
