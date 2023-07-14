package ciutils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// type TokenSetting struct {
// 	AccessTTL  int64
// 	RefreshTTL int64
// 	Passphrase string
// }

// func MakeJWTSettings(aTTL, rTTL int64, secure string) (s *TokenSetting) {
// 	s = &TokenSetting{
// 		AccessTTL:  aTTL,
// 		RefreshTTL: rTTL,
// 		Passphrase: secure,
// 	}
// 	return
// }

func CreateTokens(guid int64, passphrase string, accessTTL, refreshTTL int64) (accessToken string, refreshToken string, err error) {
	claims := &jwt.StandardClaims{
		Subject:   "access_token",
		Id:        Int64ToStr(guid),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(accessTTL)).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, err = at.SignedString([]byte(passphrase))
	if err != nil {
		return
	}

	claims.Subject = "refresh_token"
	claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(refreshTTL)).Unix()
	//rt := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	refreshToken, err = rt.SignedString([]byte(passphrase))
	if err != nil {
		accessToken = ""
	}
	fmt.Printf("CreateTokens - access_token: %v\n", accessToken)
	fmt.Printf("CreateTokens - refresh_token: %v\n", refreshToken)
	return
}
