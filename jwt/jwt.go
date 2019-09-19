package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// JWT is a Json Web Token object.
type JWT struct {
	raw string
	alg Alg

	Header *Header
	Claims Claims
}

var (
	errInvalidToken      = errors.New("invalid token")
	errInvalidKey        = errors.New("invalid key")
	errInvalidPublicKey  = errors.New("invalid public key")
	errInvalidPrivateKey = errors.New("invalid private key")
	errInvalidSign       = errors.New("invalid sign")
)

// New build an JWT entity with default value:
//     alg:    alg
//     Header: NewHeader(alg)
//     Claims: NewStdClaims()
func New(alg Alg) *JWT {
	return &JWT{
		alg:    alg,
		Header: NewHeader(alg),
		Claims: NewStdClaims(),
	}
}

// Parse read token and decode to JWT entity.
//
// The JWT entity should init header and claims for data witch to bind.
//
// Invalid token returns on error.
func (jwt *JWT) Parse(token string) error {
	jwt.raw = token
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return errInvalidToken
	}
	if part, err := base64Decode(parts[0]); err == nil {
		if err := json.Unmarshal(part, &jwt.Header); err != nil {
			return err
		}
	} else {
		return err
	}

	if part, err := base64Decode(parts[1]); err == nil {
		if err := json.Unmarshal(part, &jwt.Claims); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

// Sign for encode a token from claims by alg.
func (jwt *JWT) Sign() (string, error) {
	var token []byte
	if data, err := json.Marshal(jwt.Header); err == nil {
		token = append(token, base64Encode(data)...)
	} else {
		return "", err
	}

	if data, err := json.Marshal(jwt.Claims); err == nil {
		token = append(token, '.')
		token = append(token, base64Encode(data)...)
	} else {
		return "", err
	}

	if signData, err := jwt.alg.Sign(token); err == nil {
		token = append(token, '.')
		token = append(token, base64Encode(signData)...)
	} else {
		return "", err
	}

	jwt.raw = string(token)
	return jwt.raw, nil
}

// Verify check the token signature.
func (jwt *JWT) Verify(token string) error {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return errInvalidToken
	}
	sign, err := base64Decode(parts[2])
	if err != nil {
		return err
	}
	return jwt.alg.Verify([]byte(parts[0]+"."+parts[1]), sign)
}

func base64Decode(s string) ([]byte, error) {
	return base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(s)
}

func base64Encode(s []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(s)
}
