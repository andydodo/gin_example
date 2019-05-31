package jwt

import (
	"crypto/rsa"
	"errors"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	UserID string `json:"userID"`
	Admin  bool   `json:"admin"`
}

type JWT interface {
	GenerateToken(id string, admin bool) (string, error)
	ValidateToken(rememberHash string) (*Claims, error)
}

type DefaultJWT struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
}

func NewJWT(priKey string, pubKey string) *DefaultJWT {
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(priKey))
	if err != nil {
		log.Fatal(err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKey))
	if err != nil {
		log.Fatal(err)
	}
	return &DefaultJWT{
		signKey:   signKey,
		verifyKey: verifyKey,
	}
}

func (j *DefaultJWT) GenerateToken(id string, admin bool) (string, error) {
	c := Claims{
		UserID: id,
		Admin:  admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	tokenString, err := token.SignedString(j.signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *DefaultJWT) ValidateToken(rememberHash string) (*Claims, error) {
	c := &Claims{}
	tkn, err := jwt.ParseWithClaims(rememberHash, c, func(token *jwt.Token) (interface{}, error) {
		return j.verifyKey, nil
	})
	if err != nil {
		return &Claims{}, err
	}
	if !tkn.Valid {
		return &Claims{}, errors.New("Invalid token")
	}
	return c, nil
}
