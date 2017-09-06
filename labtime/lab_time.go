package labtime

import (
	"time"
)

const (
	OneHour = 3600
	OneDay  = 86400
)

// Get the current Time
func GetCurrentTime() int64 {
	return time.Now().Unix()
}

// Get the current time + an offset
func GetTimeAfterOffset(offset int64) int64 {
	return GetCurrentTime() + offset
}

func NewDate(day int, month time.Month, year int) int64 {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix()
}
