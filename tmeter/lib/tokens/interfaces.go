package tokens

type TokenEncoderInterface interface {
	EncodeToken(email string) (*string, error)
	DecodeToken(token string) (*string, error)
}
