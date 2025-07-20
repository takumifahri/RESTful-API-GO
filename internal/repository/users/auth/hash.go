package auth

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
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
	return b64Hash, nil
}

//  INDEVELOPMENT
// Kita byuat func enkripsi nya
func (au *authRepo) encryption(text []byte) string {
	nounce := make([]byte, 12) // Ganti dengan nounce yang sesuai
	if _, err := rand.Read(nounce); err != nil {
		return ""
	}
	encrypted := au.gcm.Seal(nil, nounce, text, nil)
	return string(encrypted)

}