package golangmodule

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type MetaToken struct {
	ID  int    `json:"id"`
	Exp string `json:"exp"`
}

type AccessToken struct {
	Claims MetaToken
}

// func Sign(Data map[string]any, expired int) (string, error) {
// 	duration, _ := strconv.Atoi(os.Getenv("JWT_TIME_DURATION"))
// 	if expired > 0 {
// 		duration = expired
// 	}

// 	drt := time.Minute * time.Duration(duration)
// 	claims := jwt.MapClaims{}
// 	claims["exp"] = time.Now().Add(drt).Unix()

// 	for i, v := range Data {
// 		claims[i] = v
// 	}
// 	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	accessToken, err := to.SignedString([]byte(os.Getenv("JWT_SECRET")))
// 	if err != nil {
// 		return "", err
// 	}

// 	return accessToken, nil
// }

func Sign(data map[string]interface{}) (string, time.Time, error) {
	// Menetapkan waktu kedaluwarsa token secara hardcode
	expiryTime := time.Now().UTC().Add(time.Hour * 48) // Waktu kedaluwarsa 48 jam di UTC

	claims := jwt.MapClaims{}
	claims["exp"] = expiryTime.Unix()

	for key, value := range data {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", time.Time{}, err
	}

	// Format waktu kedaluwarsa sebagai string untuk kemudahan penggunaan
	//expiryTimeString := expiryTime.Format(time.RFC3339)

	return accessToken, expiryTime, nil
}

func VerifyTokenHeader(requestToken string) (MetaToken, error) {

	token, err := jwt.Parse((requestToken), func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Println(err)
		return MetaToken{}, err
	}
	claimToken := DecodeToken(token)
	return claimToken.Claims, nil
}

func VerifyToken(accessToken string) (*jwt.Token, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, err := json.Marshal(&accessToken)
	if err != nil {
		return token
	}
	err = json.Unmarshal(stringify, &token)
	if err != nil {
		return token
	}
	return token
}
