package jwt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type keys struct {
	Keys []struct {
		E   string   `json:"e"`
		Kty string   `json:"kty"`
		Use string   `json:"use"`
		Kid string   `json:"kid"`
		N   string   `json:"n"`
		X5C []string `json:"x5c"`
	} `json:"keys"`
}

func Validate(token jwt.Token) (bool, error) {
	token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return lookupKey(), nil
	})

	if err == nil && token.Valid {
		fmt.Println("Authorized")
	} else {
		fmt.Println("Unauthorized")
	}

	return false, nil
}

func lookupKey() (string, error) {
	response, err := http.Get("https://api.byu.edu/.well-known/byucerts")
	if err != nil {
		return "", err
	}

	key := keys{}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(responseBody, &key)
	if err != nil {
		return "", err
	}

	return key.Keys[0].N, nil
}
