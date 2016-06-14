package wso2jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/jessemillar/jsonresp"
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
				token = token[7:] // Remove "Bearer " from the token
			} else {
				return jsonresp.New(context, http.StatusBadRequest, "No Authorization header present")
			}

			err := validate(token)
			if err != nil {
				return jsonresp.New(context, http.StatusBadRequest, err.Error())
			}

			return next(context)
		}
	}
}

func validate(token string) error {
	parsedToken, err := jwt.Parse(token, func(parsedToken *jwt.Token) (interface{}, error) {
		if parsedToken.Method.Alg() != "RS256" { // Check that our keys are signed with RS256 as expected (https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/)
			return nil, fmt.Errorf("Unexpected signing method: %v", parsedToken.Header["alg"]) // This error never gets returned to the user but may be useful for debugging/logging at some point
		}

		return lookupSigningKey(parsedToken.Header["kid"])
	})

	if parsedToken.Valid {
		return nil
	} else if validationError, ok := err.(*jwt.ValidationError); ok {
		if validationError.Errors&jwt.ValidationErrorMalformed != 0 {
			return errors.New("Authorization token is malformed")
		} else if validationError.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return errors.New("Authorization token is expired")
		}
	}

	return errors.New("Not authorized")
}

func lookupSigningKey(keyID interface{}) ([]byte, error) {
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
		if allKeys.Keys[i].Kid == keyID {
			return []byte(allKeys.Keys[0].N), nil
		}
	}

	return nil, errors.New("Could not find a valid signing key")
}
