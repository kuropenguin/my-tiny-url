package auth

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/kuropenguin/my-tiny-url/app/clock"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
	Clocker               clock.Clocker
}

type Store interface {
	Save(ctx context.Context, key string, userID int64) error
	Load(ctx context.Context, key string) (int64, error)
}

func NewJWTer(s Store) (*JWTer, error) {
	j := &JWTer{Store: s}
	privKey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	pubKey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	j.PrivateKey = privKey
	j.PublicKey = pubKey
	j.Clocker = clock.RealClocker{}
	return j, nil
}

func parse(rawKey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}
