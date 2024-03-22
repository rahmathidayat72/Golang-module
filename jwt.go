package golangmodule

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() (string, error) {
	jwtSecret, found := os.LookupEnv("JWT_SECRET")
	if !found {
		return "", errors.New("environment variable JWT_SECRET not found")
	}
	return jwtSecret, nil
}

// Fungsi untuk menghasilkan token JWT
func GenerateJWT(UserID int, expiryTime time.Time) (string, time.Time, error) {
	jwtSecret, err := getJWTSecret()
	if err != nil {
		return "", time.Time{}, err
	}
	claims := jwt.MapClaims{
		"UserID": UserID,
		"iat":    time.Now().Unix(),
		"exp":    expiryTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return strToken, expiryTime, nil
}
