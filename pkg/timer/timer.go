package timer

import (
	"fmt"
	"time"
)

// YYYY-MM-DD -> unix
// YYYY-MM-DD hh:mm:ss -> unix
func TimeString2UnixString(t string) string {
	timeTemplate := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local") // 获取时区

	if len(t) <= 10 {
		t = fmt.Sprintf("%s 00:00:00", t)
	}
	timer, _ := time.ParseInLocation(timeTemplate, t, loc)

	return fmt.Sprintf("%d", timer.Unix())
}

// YYYY-MM-DD -> unix
// YYYY-MM-DD hh:mm:ss -> unix
func TimeString2UnixInt64(t string) int64 {
	timeTemplate := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local") // 获取时区

	if len(t) <= 10 {
		t = fmt.Sprintf("%s 00:00:00", t)
	}
	timer, _ := time.ParseInLocation(timeTemplate, t, loc)

	return timer.Unix()
}
