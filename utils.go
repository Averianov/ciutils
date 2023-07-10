package utils

import (
	"ci/pkg/global"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	//"github.com/google/uuid"
	"github.com/dgrijalva/jwt-go"
)

type MessageType string

const (
	ERROR   MessageType = "Error"
	WARNING MessageType = "Warning"
	INFO    MessageType = "Info"
	SUCCESS MessageType = "Success"
)

func Message(messageType MessageType, message string) map[string]interface{} {
	return map[string]interface{}{string(messageType): message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func GetIPAddress(r *http.Request) (ipAddress string) {
	for _, ip := range strings.Split(r.RemoteAddr, ":") {
		ipAddress = ip
		if ipAddress == "[" {
			ipAddress = "127.0.0.1"
			//ipAddress = "localhost"
		}
		break
	}
	return
}

func GenConfirmCode(n int) (confirmCode string) {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("1234567890абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	confirmCode = string(b)
	return
}

func RandTo(n int) int {
	return rand.Intn(n)
}

// если err!=nil то возвращаем 0
func StrToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64) // str to int64
	if err != nil {
		return 0
	}
	return i
}

func Int64ToStr(i int64) string {
	str := fmt.Sprintf("%v", i)
	return str
}

// если err!=nil то возвращаем 0
func StrToInt(str string) int {
	i, err := strconv.Atoi(str) // str to int
	if err != nil {
		return 0
	}
	return i
}

func IntToStr(i int) string {
	str := strconv.Itoa(i)
	return str
}

func PartDateToStr(p int) string {
	var str string
	if p < 10 {
		str = "0" + strconv.Itoa(p)
	} else {
		str = strconv.Itoa(p)
	}
	return str
}

// func MonthToStr(m time.Month) string {
// 	return PartDateToStr(int(m))
// 	// var str string
// 	// if i < 10 {
// 	// 	str = "0" + strconv.Itoa(i)
// 	// } else {
// 	// 	str = strconv.Itoa(i)
// 	// }
// 	// return str
// }

// func DayToStr(d int) string {
// 	var str string
// 	if d < 10 {
// 		str = "0" + strconv.Itoa(d)
// 	} else {
// 		str = strconv.Itoa(d)
// 	}
// 	return str
// }

// func HourToStr(d int) string {
// 	var str string
// 	if d < 10 {
// 		str = "0" + strconv.Itoa(d)
// 	} else {
// 		str = strconv.Itoa(d)
// 	}
// 	return str
// }

// func MinutToStr(d int) string {
// 	var str string
// 	if d < 10 {
// 		str = "0" + strconv.Itoa(d)
// 	} else {
// 		str = strconv.Itoa(d)
// 	}
// 	return str
// }

// func SecondToStr(d int) string {
// 	var str string
// 	if d < 10 {
// 		str = "0" + strconv.Itoa(d)
// 	} else {
// 		str = strconv.Itoa(d)
// 	}
// 	return str
// }

// Возвращает строку времени в формате: "2006-01-02 15:04:05"
func GetTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

// Возвращает строку времени в формате: "20060102150405"
func GetShortTime() string {
	st := GetTime()
	st = strings.Replace(st, " ", "", -1)
	st = strings.Replace(st, "-", "", -1)
	st = strings.Replace(st, ":", "", -1)
	fmt.Println(st)
	return st
}

func CreateTokens(guid int64) (accessToken string, refreshToken string, err error) {
	claims := &jwt.StandardClaims{
		Subject:   "access_token",
		Id:        Int64ToStr(guid),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(global.Config.TTLAccessToken)).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, err = at.SignedString([]byte(global.Config.Passphrase))
	if err != nil {
		return
	}

	claims.Subject = "refresh_token"
	claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(global.Config.TTLRefreshToken)).Unix()
	//rt := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	refreshToken, err = rt.SignedString([]byte(global.Config.Passphrase))
	if err != nil {
		accessToken = ""
	}
	fmt.Printf("CreateTokens - access_token: %v\n", accessToken)
	fmt.Printf("CreateTokens - refresh_token: %v\n", refreshToken)
	return
}
