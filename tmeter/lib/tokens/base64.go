package tokens

import (
	"encoding/base64"
	"errors"
)

type base64Encoder struct{}

func (s *base64Encoder) EncodeToken(text string) (*string, error) {
	if text == "" {
		return nil, errors.New("email is empty")
	}
	str := base64.StdEncoding.EncodeToString([]byte(text))
	return &str, nil
}

func (s *base64Encoder) DecodeToken(token string) (*string, error) {
	b, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	v := string(b)
	return &v, nil
}

func NewBase64TokenEncoder() TokenEncoderInterface {
	return &base64Encoder{}
}
