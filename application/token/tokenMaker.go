package token

import (
	"time"

	"github.com/garcia-paulo/go-gin/infra/config"
	"github.com/o1egl/paseto"
)

type TokenMaker struct {
	token        *paseto.V2
	symmetricKey []byte
}

func NewTokenMaker(config *config.Config) *TokenMaker {
	maker := &TokenMaker{
		token:        paseto.NewV2(),
		symmetricKey: []byte(config.TokenKey),
	}
	return maker
}

func (m *TokenMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return m.token.Encrypt(m.symmetricKey, payload, nil)
}

func (m *TokenMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := m.token.Decrypt(token, m.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Validate()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
