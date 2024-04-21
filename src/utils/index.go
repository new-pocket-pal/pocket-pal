package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"pocket-pal/src/domain/entity"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateSalt() string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(salt)
}

// HashPassword to hash password
func HashPassword(password, salt string) string {
	// implementation of HashPassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return string(hashedPassword)
}

func GenerateAccessToken(user *entity.User, secret string, expireTime int64) (string, error) {
	// implementation of GenerateAccessToken
	expirationTime := time.Now().Add(time.Duration(expireTime) * time.Second)

	// create claims
	claims := &entity.Claims{
		UserID: user.UserID,
		Phone:  user.Phone,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return accessToken, nil
}

func DecodeJwtToken(tokenString string, secret string) (*entity.Claims, error) {
	// implementation of DecodeJwtToken
	// parse token
	token, err := jwt.ParseWithClaims(tokenString, &entity.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// get claims
	claims, ok := token.Claims.(*entity.Claims)
	if !ok {
		log.Fatal(err)
		return nil, err
	}

	return claims, nil
}

func ComparePassword(password, hashedPassword string) bool {
	// implementation of CompainPassword
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
