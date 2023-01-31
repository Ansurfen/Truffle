package utils

import (
	"strconv"
	"time"
)

func NowTimestamp() int64 {
	return time.Now().Unix()
}

func NowTimestampByString() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

func CmpTimestamp(a, b, offset int64) int {
	if a+offset == b {
		return 0
	}
	if a+offset < b {
		return -1
	}
	return 1
}
