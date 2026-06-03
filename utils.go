package ciutils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type MessageType string

const (
	ERROR   MessageType = "Error"
	WARNING MessageType = "Warning"
	INFO    MessageType = "Info"
	SUCCESS MessageType = "Success"

	LOCALHOST string = "127.0.0.1"
	RUNES     string = "1234567890абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"

	DB   string = "db"
	JSON string = "json"
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
		if len(strings.Split(ip, ".")) == 4 {
			ipAddress = ip
		}
		if ipAddress == "[" {
			ipAddress = LOCALHOST
			break
		}
	}
	return
}

func GenConfirmCode(n int) (confirmCode string) {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune(RUNES)

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
func StrToFloat64(str string) float64 {
	i, err := strconv.ParseFloat(str, 64) // str to float64
	if err != nil {
		return 0
	}
	return i
}

// если err!=nil то возвращаем 0
func StrToInt(str string) int {
	i, err := strconv.Atoi(str) // str to int
	if err != nil {
		return 0
	}
	return i
}

func Int64ToStr(i int64) string {
	str := fmt.Sprintf("%v", i)
	return str
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

func GetPathSeparator() (separator string) {
	if runtime.GOOS == "windows" {
		separator = "\\"
	} else {
		separator = "/"
	}
	return
}

func MakeSureFileExists(fullFileName string) (file *os.File, err error) {
	if _, err = os.Stat(fullFileName); os.IsNotExist(err) {
		if file, err = os.Create(fullFileName); os.IsNotExist(err) {
			//if err = os.Mkdir(fullFileName, 0777); os.IsNotExist(err) {
			return
		}
	}
	return
}

func BoolToStr(incoming bool) string {
	if incoming {
		return "true"
	}
	return "false"
}

func StrToBool(incoming string) bool {
	if incoming == "true" || incoming == "1" {
		return true
	}
	return false
}
