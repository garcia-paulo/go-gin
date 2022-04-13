package token

import (
	"time"

	"github.com/garcia-paulo/go-gin/infra/config"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(config *config.Config) *PasetoMaker {
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(config.TokenKey),
	}
	return maker
}

func (m *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return m.paseto.Encrypt(m.symmetricKey, payload, nil)
}

func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := m.paseto.Decrypt(token, m.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Validate()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
