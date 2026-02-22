package utility

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// password comparision at the time of login
func CheckPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateJWT(userID string, email string, jwtSecret string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseJWT(tokenString string, jwtSecret string) (string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return "", err
	}

	// extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("my claims", claims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	userID, ok := claims["user_id"].(string)
	fmt.Println("my okkk valiue", ok)
	if ok == false {
		return "", errors.New("invalid token claims")
	}
	return userID, nil
}
