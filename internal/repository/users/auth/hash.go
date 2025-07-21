package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	cryptFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
)

// kita akan implemen user hash disini
func (au *authRepo) GenerateUserHash(password string) (hash string, err error) {
	// /Kita akan menggunaakn argon2 utnuk encrypt password mari kita coba intip
	salt := make([]byte, 16) // Ganti dengan salt yang sesuai
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	argonHash := argon2.IDKey([]byte(password), salt, au.time, au.memory, au.threads, au.keylen)

	b64Hash := au.encryption(argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf(cryptFormat, argon2.Version, au.memory, au.time, au.threads, b64Salt, b64Hash)
	return encodedHash, nil
}

//  INDEVELOPMENT
// Kita byuat func enkripsi nya
func (au *authRepo) encryption(text []byte) string {
	nonce := make([]byte, au.gcm.NonceSize())

	cipherText := au.gcm.Seal(nonce, nonce, text, nil )
	return base64.StdEncoding.EncodeToString(cipherText)
}