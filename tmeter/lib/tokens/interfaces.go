package tokens

type TokenEncoderInterface interface {
	EncodeToken(text string) (*string, error)
	DecodeToken(token string) (*string, error)
}
