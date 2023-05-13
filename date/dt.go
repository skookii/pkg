package date

import (
	"strconv"
	"time"
)

func Today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

func Yesterday() time.Time {
	now := time.Now().AddDate(0, 0, -1)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

func TimeToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func Tomorrow() time.Time {
	now := time.Now().AddDate(0, 0, 1)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

// time to unix string
func UnixTimeString(tm time.Time) string {
	i64 := tm.Unix()
	return strconv.FormatInt(i64, 10)
}

// unix string to time
func StringToTime(s string) time.Time {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local)
	}
	return time.Unix(i64, 0)
}
