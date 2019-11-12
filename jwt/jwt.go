package jwt

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/lestrrat-go/jwx/jwk"
)

var jwksEndpoint string

func stripBearerPrefix(token string) string {
	if strings.HasPrefix(strings.ToLower(token), "bearer") {
		return token[len("Bearer "):]
	}
	return token
}

//VerifyToken verifies the signature against the given endpoint
func VerifyToken(token, jwks string) error {
	jwksEndpoint = jwks
	_, err := jwt.Parse(stripBearerPrefix(token), getKey)
	if err != nil {
		return fmt.Errorf("could not parse jwt: %v", err)
	}
	return nil
}

func getKey(token *jwt.Token) (interface{}, error) {
	set, err := jwk.FetchHTTP(jwksEndpoint)
	if err != nil {
		return nil, fmt.Errorf("could not fetch key set: %v", err)
	}
	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("kid not found in header")
	}

	if key := set.LookupKeyID(kid); len(key) == 1 {
		return key[0].Materialize()
	}
	return nil, fmt.Errorf("unable to find key %q", kid)
}
