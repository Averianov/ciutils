package ciutils

import (
	"time"
)

type TimeCounter struct {
	T int64
}

func StartTimeCounter() TimeCounter {
	return TimeCounter{T: time.Now().UnixNano() / int64(time.Second)}
}

func (t TimeCounter) StopTimeCounter() int64 {
	return time.Now().UnixNano()/int64(time.Second) - t.T
}
