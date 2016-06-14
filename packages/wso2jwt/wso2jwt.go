package wso2jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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

// ValidateJWT is the middleware function
func ValidateJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			token := context.Request().Header().Get("Authorization")
			if token != "" {
				token = token[7:]
			} else {
				return echo.NewHTTPError(http.StatusBadRequest, "No Authorization header present")
			}

			fmt.Println(token)

			valid, err := validate(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			fmt.Println("Hit this line")

			if valid == false {
				fmt.Println("Returning unauthorized error")
				return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
			}

			return nil
		}
	}
}

func validate(headerToken string) (bool, error) {
	token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		return lookupKey(token.Header["kid"])
	})

	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return false, errors.New("Timing is everything")
		} else {
			return false, errors.New("Couldn't handle the token: " + err.Error())
		}
	} else {
		return false, errors.New("Couldn't handle the token: " + err.Error())
	}
}

func lookupKey(kid interface{}) (string, error) {
	response, err := http.Get("https://api.byu.edu/.well-known/byucerts")
	if err != nil {
		return "", err
	}

	allKeys := keys{}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(responseBody, &allKeys)
	if err != nil {
		return "", err
	}

	for i := range allKeys.Keys {
		if allKeys.Keys[i].Kid == kid {
			return allKeys.Keys[0].N, nil
		}
	}

	return "", errors.New("Could not find valid key")
}
