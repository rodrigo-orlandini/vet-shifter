package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTPayload struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

type AuthClaims struct {
	jwt.RegisteredClaims
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

func IssueJWT(userID string, email string, userType string, rememberMe bool) (string, time.Time, error) {
	secret := GetJWTSecret()
	if secret == "" {
		return "", time.Time{}, fmt.Errorf("JWT_SECRET is not set")
	}

	var expDuration time.Duration
	if rememberMe {
		expDuration = 7 * 24 * time.Hour
	} else {
		expDuration = 24 * time.Hour
	}

	now := time.Now()
	exp := now.Add(expDuration)

	claims := AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		Sub:   userID,
		Email: email,
		Type:  userType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", time.Time{}, err
	}

	return signed, exp, nil
}

func VerifyJWT(tokenString string) (*AuthClaims, error) {
	secret := GetJWTSecret()
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET is not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AuthClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func JWTPayloadFromClaims(claims *AuthClaims) JWTPayload {
	return JWTPayload{
		ID:    claims.Sub,
		Email: claims.Email,
		Type:  claims.Type,
	}
}
