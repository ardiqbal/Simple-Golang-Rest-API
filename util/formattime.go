package util

import "time"

func Format(utime string) string {
	timeLayout := "2006-01-02 15:04:05"
	t, _ := time.Parse(timeLayout, utime)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	var t2 = t.In(loc)
	var t3 = t2.Format("2006-01-02T15:04:05.999999Z0700")
	return t3
}
