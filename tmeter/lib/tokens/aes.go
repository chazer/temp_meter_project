package tokens

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

func encrypt(t string, key []byte) (string, error) {
	text := []byte(t)
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, text, nil)), nil
}

func decrypt(cryptogram string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cryptogram)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", err
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

type aesEncoder struct {
	secret []byte
}

func (s *aesEncoder) EncodeToken(data string) (*string, error) {
	if data == "" {
		return nil, errors.New("email is empty")
	}
	str, err := encrypt(data, s.secret)
	if err != nil {
		return nil, errors.New("unsupported value")
	}
	return &str, nil
}

func (s *aesEncoder) DecodeToken(token string) (*string, error) {
	str, err := decrypt(token, s.secret)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	return &str, nil
}

func NewAesTokenEncoder(secret string) TokenEncoderInterface {
	hasher := md5.New()
	hasher.Write([]byte(secret))
	secret = hex.EncodeToString(hasher.Sum(nil))
	// secret must be 32 bytes always
	return &aesEncoder{
		secret: []byte(secret),
	}
}
