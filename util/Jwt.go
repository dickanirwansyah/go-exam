package util

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey []byte

type DataUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func init() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file !")
	}

	readSecretKey := os.Getenv("JSON_WEB_TOKEN_SECRET_KEY")
	if readSecretKey == "" {
		panic("JWT secret key is not set the environment variables !")
	}

	jwtSecretKey = []byte(readSecretKey)
}

func GenerateJwt(dataUser DataUser) (string, error) {

	//create new token & token expiration 1 hours
	claims := jwt.MapClaims{
		"data_user": dataUser,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}

	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return createToken.SignedString(jwtSecretKey)
}

func CheckHashBcryptPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func VerifyJwt(jwtToken string) (DataUser, jwt.MapClaims, error) {

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, success := token.Method.(*jwt.SigningMethodHMAC); !success {
			return nil, fmt.Errorf("Unexpected signing method : %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return DataUser{}, nil, err
	}

	// extract claims if tokens is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		dataUserJSON, ok := claims["data_user"].(map[string]interface{})
		if !ok {
			return DataUser{}, nil, fmt.Errorf("Invalid json data_user claim")
		}

		//process marshall data_user to JSON  and then unmarshall it into the struct
		dataUserBytes, err := json.Marshal(dataUserJSON)
		if err != nil {
			return DataUser{}, nil, fmt.Errorf("Error marshall data_user : %v", err)
		}

		var dataUser DataUser
		err = json.Unmarshal(dataUserBytes, &dataUser)
		if err != nil {
			return DataUser{}, nil, fmt.Errorf("Error unmarshall data_uer : %v", err)
		}

		//return the extract data_user and other claims
		return dataUser, claims, nil
	}

	//if not ok
	return DataUser{}, nil, fmt.Errorf("Invalid token")
}
