package provider

type TokenProvider interface {
	GenerateToken(data interface{}) (string, error)
	VerifyToken(token string) (interface{}, error)
}
