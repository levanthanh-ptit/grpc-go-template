package provider

type TokenProvider interface {
	GenerateToken() (string, error)
	VerifyToken(token string) (interface{}, error)
}
