package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SecretKey = "SECRET_KEY"

type CustomClaim struct {
	Username             string
	Roles                []string
	jwt.RegisteredClaims // Registerd claims are predefined fields
}

func main() {
	// 1. Create custom claim with predefined claims (RegisteredClaims)
	claims := CustomClaim{
		Username: "john",
		Roles:    []string{"FullAccess", "ReadWrite"},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MyApp",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// 2. Create tokem witn claims with specific signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. Create jwt token sign secret key
	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// ========

	// Validate token
	token, err = jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, valid := t.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		panic(err)
	}

	if token.Valid {
		fmt.Println("TOKEN VALID")
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
	} else {
		panic("token invalid")
	}
}
