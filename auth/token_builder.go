package auth

import (
	"encoding/json"
	"time"

	"aidanwoods.dev/go-paseto"
)

type TokenBuilder interface {
	Build(user User) (token string, err error)
}

func NewPasetoBuilder(secretKey string) (TokenBuilder, error) {
	key, err := paseto.V4SymmetricKeyFromHex(secretKey)
	if err != nil {
		return nil, err
	}
	return &pasetoBuilder{
		secretKey: key,
	}, nil
}

type pasetoBuilder struct {
	secretKey paseto.V4SymmetricKey
}

func (b *pasetoBuilder) Build(user User) (token string, err error) {
	now := time.Now()

	tok := paseto.NewToken()
	tok.SetIssuedAt(now)
	tok.SetNotBefore(now)
	tok.SetExpiration(now.Add(2 * time.Hour))
	tok.SetIssuer("com.mcosta74.auth")

	userStr, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	tok.SetString("userInfo", string(userStr))

	return tok.V4Encrypt(b.secretKey, nil), nil
}
