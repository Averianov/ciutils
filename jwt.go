package ciutils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenSetting struct {
	AccessTTL  int64
	RefreshTTL int64
	Passphrase string
}

func MakeJWTSettings(aTTL, rTTL int64, secure string) (s *TokenSetting) {
	s = &TokenSetting{
		AccessTTL:  aTTL,
		RefreshTTL: rTTL,
		Passphrase: secure,
	}
	return
}

func (s *TokenSetting) CreateTokens(guid int64) (accessToken string, refreshToken string, err error) {
	claims := &jwt.StandardClaims{
		Subject:   "access_token",
		Id:        Int64ToStr(guid),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(s.AccessTTL)).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, err = at.SignedString([]byte(s.Passphrase))
	if err != nil {
		return
	}

	claims.Subject = "refresh_token"
	claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(s.RefreshTTL)).Unix()
	//rt := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	refreshToken, err = rt.SignedString([]byte(s.Passphrase))
	if err != nil {
		accessToken = ""
	}
	fmt.Printf("CreateTokens - access_token: %v\n", accessToken)
	fmt.Printf("CreateTokens - refresh_token: %v\n", refreshToken)
	return
}
