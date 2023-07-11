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
	//"github.com/google/uuid"
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
		ipAddress = ip
		if ipAddress == "[" {
			ipAddress = LOCALHOST
		}
		break
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

func GetPathSeparator() (separator string) {
	if runtime.GOOS == "windows" {
		separator = "\\"
	} else {
		separator = "/"
	}
	return
}

func GetDBPath(rawToLog bool) (dir string, err error) {
	dir, err = GetRootDir()
	dir = dir + GetPathSeparator()
	if rawToLog {
		dir = dir + "json"
	} else {
		dir = dir + "db"
	}
	return
}

func GetRootDir() (rootDir string, err error) {
	rootDir, err = os.Getwd()
	if err != nil {
		return
	}
	if _, err = os.Stat(rootDir + GetPathSeparator() + "db"); os.IsNotExist(err) {
		fmt.Printf("## SetRootDir - Stat err: %s\n", err.Error())
		if err = os.Mkdir(rootDir+GetPathSeparator()+"db", 0777); os.IsNotExist(err) {
			fmt.Printf("## SetRootDir - Mkdir err: %s\n", err.Error())
			return
		}
	}
	fmt.Printf("RootDir: %s\n", rootDir)
	return
}

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
