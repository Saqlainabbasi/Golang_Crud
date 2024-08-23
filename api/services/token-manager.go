package services

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// ...interface
type TokenManager interface {
	NewJWT(userId int64) (string, error)
	Parse(token string) (int64, error)
	NewRefreshToken() (string, error)
}

// ...struct
type tokenManager struct {
	signingKey string
}

// ...constructor func
func NewTokenManager(signedKey string) TokenManager {
	return &tokenManager{signingKey: signedKey}
}

// ...implementation of interface
// this a receiver func, receiver t is of type tokenManager struct..
// and this struct implements the interface...
// Subject:   strconv.FormatInt(userId,10),
// this will convert the int64 to string
func (t *tokenManager) NewJWT(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 120).Unix(),
		Subject:   strconv.FormatInt(userId, 10),
	})
	return token.SignedString([]byte(t.signingKey))
}

// ...func to verify the token
// Parse parses the given access token and returns the user ID associated with it.
// It takes the access token as a parameter and returns the user ID as an int64 and an error.
// If the access token is invalid or the user ID cannot be retrieved from the token, an error is returned.
func (t *tokenManager) Parse(accessToken string) (int64, error) {
	// fmt.Println(accessToken)
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method}")
		}
		return []byte(t.signingKey), nil

	})
	if err != nil {
		return 0, err

	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("cannot get claims form the token")
	}
	atoi, err := strconv.Atoi(claims["sub"].(string))
	if err != nil {
		return 0, errors.New("conversion failed")
	}
	id := int64(atoi)
	return id, nil
}

// .....func for refresh token
func (t *tokenManager) NewRefreshToken() (string, error) {
	// make a slice of byte of length 32
	b := make([]byte, 32)

	// declare variable to hold the source(in this case second) value generated form the rand func
	s := rand.NewSource(time.Now().Unix())
	// random no form source of seconds...
	r := rand.New(s)

	//
	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return string(b), nil

}
