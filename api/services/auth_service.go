package services

import (
	"am/office-check-in/database"
	"am/office-check-in/models"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func Register(user models.UserBody) (models.User, error) {
	user.Password = HashPassword(user.Password)

	createdUser, err := CreateUser(user)

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

type JwtClaims struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	jwt.RegisteredClaims
}

func Login(email string, password string) (string, error) {
	dbConnection := database.Connection()
	var user models.User

	dbConnection.Model(&models.User{}).Where("email = ?", email).First(&user)

	if !(HashPassword(password) == user.Password) {
		return "", errors.New("invalid credentials")
	}

	claims := &JwtClaims{
		Name:    user.Name,
		Picture: user.Picture,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		panic("JWT_SECRET is not set")
	}

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
