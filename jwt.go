package ciutils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	sl "github.com/Averianov/cisystemlog"
)

func init() {
	if sl.L == nil {
		sl.CreateLogs(jwt, "./log/", 3, 0)
	}
}

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
	sl.L.Debug("CreateTokens - access_token: %v\n", accessToken)
	sl.L.Debug("CreateTokens - refresh_token: %v\n", refreshToken)
	return
}

func CheckToken(r *http.Request, passphrase string) (guid int64, incomingToken string, err error) {
	var jwToken *jwt.Token

	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		err = fmt.Errorf("%s", "Check token - token not founded")
		return
	}
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		err = fmt.Errorf("%s", "Check token - token is not full")
		return
	}
	incomingToken = splitted[1]
	tk := &jwt.StandardClaims{}
	jwToken, err = jwt.ParseWithClaims(incomingToken, tk, func(jwToken *jwt.Token) (interface{}, error) {
		return []byte(passphrase), nil
	})
	if err != nil {
		return
	}
	sl.L.Debug("jwToken: %v\n", jwToken)
	sl.L.Debug("tk: %v\n", tk)
	if !jwToken.Valid {
		err = fmt.Errorf("%s", "Check token - token not valid")
		return
	}
	guid = StrToInt64(tk.Id)
	if guid == 0 {
		err = fmt.Errorf("%s", "Check token - user number not founded in token")
		return
	}
	// _, err = uuid.Parse(guid)
	// if err != nil {
	// 	sl.L.Debug("uuid.Parse error: %v\n", err)
	// }
	return
}
