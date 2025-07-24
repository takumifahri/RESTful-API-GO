package auth

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
	"github.com/takumifahri/RESTful-API-GO/internal/tracing"
)

const (
	cryptFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
)

// kita akan implemen user hash disini
func (au *authRepo) GenerateUserHash(ctx context.Context,password string) (hash string, err error) {
	_, span := tracing.CreateSpanWrapper(ctx, "GenerateUserHash")
	defer span.End()
	// /Kita akan menggunaakn argon2 utnuk encrypt password mari kita coba intip
	salt := make([]byte, 16) // Ganti dengan salt yang sesuai
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	argonHash := argon2.IDKey([]byte(password), salt, au.time, au.memory, au.threads, au.keylen)

	b64Hash := au.encryption(ctx, argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf(cryptFormat, argon2.Version, au.memory, au.time, au.threads, b64Salt, b64Hash)
	return encodedHash, nil
}

// Kita byuat func enkripsi nya
func (au *authRepo) encryption(ctx context.Context, text []byte) string {
	_, span := tracing.CreateSpanWrapper(ctx, "encryption")
	defer span.End()
	nonce := make([]byte, au.gcm.NonceSize())

	cipherText := au.gcm.Seal(nonce, nonce, text, nil)
	return base64.RawStdEncoding.EncodeToString(cipherText)
}

func (au *authRepo) decrypt(ctx context.Context, cipherText string) ([]byte, error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "decrypt")
	defer span.End()
	decode, err := base64.RawStdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, fmt.Errorf("failed to decode cipher text: %v", err)
	}

	// bandingkan
	if len(decode) < au.gcm.NonceSize() {
		return nil, fmt.Errorf("cipher text too short")
	}

	return au.gcm.Open(nil,
		decode[:au.gcm.NonceSize()],
		decode[au.gcm.NonceSize():],
		nil,
	)
}

func (au *authRepo) comparePassword(ctx context.Context, password, hash string) (bool,error) {
	ctx, span := tracing.CreateSpanWrapper(ctx, "comparePassword")
	defer span.End()
	parts := strings.Split(hash, "$") // ini akan membagi hash menjadi beberapa bagian

	// variable kosong 
	var memory, time uint32
	var parallesim uint8

	switch parts[1] {
		case "argon2id":
			_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &parallesim)
			if err != nil {
				return false, fmt.Errorf("failed to parse hash: %v", err)
			}
			salt, err := base64.RawStdEncoding.DecodeString(parts[4])
			if err != nil {
				return false, fmt.Errorf("failed to decode salt: %v", err)
			}

			hash := parts[5]

			decryptHash, err := au.decrypt(ctx, hash)
			if err != nil {
				return false, fmt.Errorf("failed to decrypt hash: %v", err)
			}

			// Kita buat keylen var
			var keyLen = uint32(len(decryptHash))
			// kita compare skearang
			compareHash := argon2.IDKey([]byte(password), salt, time, memory, parallesim, keyLen)

			return subtle.ConstantTimeCompare(compareHash, decryptHash) == 1, nil
	}
	return false, nil
}