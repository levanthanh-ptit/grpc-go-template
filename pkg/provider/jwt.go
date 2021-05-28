package provider

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type AddClaimsFunc = func(target jwt.MapClaims, data interface{})

type ExtractClaimsFunc = func(data jwt.Claims) interface{}

type JWTProvider struct {
	// Signing part
	signingKey           string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
	addClaims            AddClaimsFunc
	extractClaims        ExtractClaimsFunc
	// Persistance part
	// TODO: store meta with redit
}

func NewJWTProvider(
	signingKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
	makeClaims AddClaimsFunc,
	extractClaims ExtractClaimsFunc,
) *JWTProvider {
	return &JWTProvider{
		signingKey,
		accessTokenDuration,
		refreshTokenDuration,
		makeClaims,
		extractClaims,
	}
}

func (p *JWTProvider) GenerateToken(data interface{}) (string, error) {
	claims := jwt.MapClaims{}
	p.addClaims(claims, data)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(p.accessTokenDuration).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(p.signingKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (p *JWTProvider) VerifyToken(token string) (interface{}, error) {
	getKey := func(token *jwt.Token) (interface{}, error) {
		return []byte(p.signingKey), nil
	}
	jwtToken, err := jwt.Parse(token, getKey)
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, errors.New("TOKEN_INVALID")
	}
	data := p.extractClaims(jwtToken.Claims)
	return data, nil
}
