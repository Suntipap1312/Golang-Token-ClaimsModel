package main

import (
	"fmt"
	"time"
)

func main() {
	type JwtClaims struct {
		Username string `json:"username",omitempty`
		Password string `json:"password",omitempty`
		jwt.StandardClaims
	}
}

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(ip, true) {
		return nil
	}
	return fmt.Errorf("token is invalid")
}
