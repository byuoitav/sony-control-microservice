package jwt

import (
	"encoding/json"
	"errors"
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

func Validate(myToken string, myLookupKey func(interface{}) (interface{}, error)) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return lookupKey(token.Header["kid"])
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

	return nil, nil
}

func lookupKey(kid interface{}) (interface{}, error) {
	response, err := http.Get("https://api.byu.edu/.well-known/byucerts")
	if err != nil {
		return nil, err
	}

	allKeys := keys{}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &allKeys)
	if err != nil {
		return nil, err
	}

	for i := range allKeys.Keys {
		if allKeys.Keys[i].Kid == kid {
			return allKeys.Keys[0].N, nil
		}
	}

	return nil, errors.New("Could not find valid key")
}
