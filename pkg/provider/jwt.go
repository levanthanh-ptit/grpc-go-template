package provider

type JWTProvider struct {
}

func NewJWTProvider(signingKey string) *JWTProvider {
	return &JWTProvider{}
}

func (p *JWTProvider) GenerateToken() (string, error) {
	return "", nil
}

func (p *JWTProvider) VerifyToken(token string) (interface{}, error) {
	return nil, nil
}
