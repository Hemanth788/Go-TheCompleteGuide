package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "secret"

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userID": userID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(SECRET_KEY)) // signing with a secret, though the param is of type any, you must pass a byte slice
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // checking the signing method type

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(SECRET_KEY), nil // again, byte slice only, even with any in the return type
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return  0, errors.New("Invalid token") // checking if the parsed token is valid, can be invalid if exp
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims) // checking claims type 
	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	// email := claims["email"].(string)
	userID := int64(claims["userID"].(float64)) // accessing claims, though I set int66(see L12, L15), it is stored as float64, so checking for than and converting it into int64 for my use case
	return userID, nil
}