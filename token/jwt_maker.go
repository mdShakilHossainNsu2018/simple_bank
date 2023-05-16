package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, ErrInvalidSecretKey
	}

	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	// Create a new token object, specifying signing method and the payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Create a string from the token object
	return token.SignedString([]byte(maker.secretKey))
}

// VerifyToken verifies the token string and return a token object if the token string is valid
func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	// Parse the token string and store the result in the token object
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		// Return the secret key used to create the token
		return []byte(maker.secretKey), nil
	})

	// Check if there is any error
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	payload, ok := token.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
