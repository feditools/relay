package models

import (
	"fmt"
	"github.com/feditools/relay/internal/config"
	"github.com/spf13/viper"
	"github.com/tmthrgd/go-hex"
	"testing"
)

func TestDecrypt(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	byts := hex.MustDecodeString("43dc49ab017fbde685011bc75e7aeecf46e2e6ca2d960681ebca6b7d9b5a74ad0348cfcadbdb71bebb")

	val, err := decrypt(byts)
	if err != nil {
		t.Errorf("unexpected error getting scanning, got: '%s', want: 'nil", err)
		return
	}
	if string(val) != "test string 1" {
		t.Errorf("unexpected value, got: '%s', want: '%s'", val, "test string 1")
	}
}

func TestDecrypt_NoData(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	var byts []byte

	_, err := decrypt(byts)
	errMsg := "data too small"
	if err == nil {
		t.Errorf("expected error getting scanning, got: 'nil', want: '%s", errMsg)
		return
	}
	if err.Error() != errMsg {
		t.Errorf("unexpected error getting scanning, got: '%s', want: '%s", err.Error(), errMsg)
		return
	}
}

func TestEncrypt(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	tables := []struct {
		n string
	}{
		{"test string 1"},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Getting id", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			data, err := encrypt([]byte(table.n))
			if err != nil {
				t.Errorf("unexpected error getting value: %s", err.Error())
				return
			}

			gcm, err := getCrypto()
			if err != nil {
				t.Errorf("getting crypto: %s", err.Error())
				return
			}

			nonceSize := gcm.NonceSize()
			if len(data) < nonceSize {
				t.Errorf("value too small")
				return
			}

			nonce, ciphertext := data[:nonceSize], data[nonceSize:]
			plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
			if err != nil {
				t.Errorf("decrypting: %s", err.Error())
				return
			}
			if string(plaintext) != table.n {
				t.Errorf("unexpected value, got: '%s', want: '%s'", plaintext, table.n)
			}
		})
	}
}

func BenchmarkDecrypt(b *testing.B) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	byts := hex.MustDecodeString("43dc49ab017fbde685011bc75e7aeecf46e2e6ca2d960681ebca6b7d9b5a74ad0348cfcadbdb71bebb")

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = decrypt(byts)
		}
	})
}

func BenchmarkEncrypt(b *testing.B) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	byts := []byte("test string")

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = encrypt(byts)
		}
	})
}
