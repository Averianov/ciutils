package ciutils

import (
	"strings"
	"time"
)

const (
	timeLayout            string = "02.01.2006 15:04:05 MST -07:00"
	DEFAULT_NAME_LOCATION string = "UTC"
)

var TimeLocation *time.Location

func init() {
	TimeLocation, _ = time.LoadLocation(DEFAULT_NAME_LOCATION)
}

func Now() (t time.Time) {
	t = time.Now().In(TimeLocation)
	return
}

// Return simple string time in format: "02.01.2006 15:04:05 MST -07:00"
func TimeToString(t time.Time) string {
	return t.Format(timeLayout)
}

// Return simple string time in arbitrary format
func TimeToStringInFormat(t time.Time, layout string) string {
	return t.Format(layout)
}

// Return simple string time in format: "2006-01-02 15:04:05"
func GetTime() string {
	t := Now()
	return t.Format("2006-01-02 15:04:05")
}

// Return short string time in format: "20060102150405"
func GetShortTime() string {
	st := TimeToStringInFormat(Now(), "20060102150405")
	// st := GetTime()
	// st = strings.Replace(st, " ", "", -1)
	// st = strings.Replace(st, "-", "", -1)
	// st = strings.Replace(st, ":", "", -1)
	// fmt.Println(st)
	return st
}

func ParseStringToTime(t, layout string) (result time.Time, err error) {
	t = strings.TrimSpace(t)
	result, err = time.Parse(layout, t)
	if err != nil {
		return
	}
	result = result.In(TimeLocation)
	return
}

func TimeToInt64(t time.Time) int64 {
	return t.UnixMilli()
}

func Int64ToTime(i int64) time.Time {
	return time.UnixMilli(i).In(TimeLocation)
}

func SecondToInt64(i int) int64 {
	return int64(time.Second/time.Millisecond) * int64(i)
}

func TimeToInt64(t time.Time) int64 {
	return t.UnixMilli()
}

func Int64ToTime(i int64) time.Time {
	return time.UnixMilli(i).In(TimeLocation)
}

func SecondToInt64(i int) int64 {
	return int64(time.Second/time.Millisecond) * int64(i)
}

func GetNumberOfDay(t time.Time) (dayNumber int64) {
	initialСount := Int64ToTime(0)
	dayNumber = int64(t.Sub(initialСount).Hours() / 24)
	return
}

func GetDay(t time.Time) (day int64) {
	day = int64(t.Day())
	return
}

func GetMonth(t time.Time) (month int64) {
	month = int64(t.Month())
	return
}

func GetTomorrowTime(t time.Time) (tomorrow time.Time) {
	tomorrow = t.AddDate(0, 0, +1).In(TimeLocation)
	return
}

func GetYesterdayTime(t time.Time) (yesterday time.Time) {
	yesterday = t.AddDate(0, 0, -1).In(TimeLocation)
	return
}

func GetPreviousMonthTime(t time.Time) (yesterday time.Time) {
	yesterday = t.AddDate(0, -1, 0).In(TimeLocation)
	return
}

func GetPreviousHoursTime(t time.Time, count time.Duration) (previousTime time.Time) {
	d := time.Duration(-(time.Hour * count))
	previousTime = t.Add(d).In(TimeLocation)
	return
}

// Get period from first day to current day of current month
func GetCurrentPeriod(now time.Time) (fromDate time.Time, toDate time.Time) {
	fromDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 000, TimeLocation)
	toDate = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999, TimeLocation)
	return
}

func IsFirstWorkingDay(created, now time.Time, location *time.Location) bool {
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 000, location)
	return now.Sub(created) < 0
}

// Get period from first day to last day of last month
func GetLastPeriod(now time.Time) (fromDate time.Time, toDate time.Time) {
	previousMonth := GetPreviousMonthTime(now)
	// first day and time in last month
	fromDate = time.Date(previousMonth.Year(), previousMonth.Month(), 1, 0, 0, 0, 000, TimeLocation)
	// first day in current month
	firstDateInCurrentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 000, TimeLocation)
	// last day in last month
	lastDateInPreviousMonth := GetYesterdayTime(firstDateInCurrentMonth)
	// last day and time in last month
	toDate = time.Date(lastDateInPreviousMonth.Year(), lastDateInPreviousMonth.Month(),
		lastDateInPreviousMonth.Day(), 23, 59, 59, 999, TimeLocation)

	//fromDate = GetTomorrowTime(fromDate)
	//toDate = GetTomorrowTime(toDate)
	return
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
