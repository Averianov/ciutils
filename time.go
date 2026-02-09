package ciutils

import (
	"strings"
	"time"
)

const (
	timeLayout = "02.01.2006 15:04:05 MST -07:00"
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

func TimeToString(t time.Time) string {
	return t.Format(timeLayout)
}

func TimeToStringInFormat(t time.Time, layout string) string {
	return t.Format(layout)
}

func ParseStringToTime(t, layout string) (result time.Time, err error) {
	t = strings.TrimSpace(t)
	result, err = time.Parse(layout, t)
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
