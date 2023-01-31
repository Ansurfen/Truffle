package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandValue(args ...any) string {
	res := ""
	for _, v := range args {
		res += ToString(v)
	}
	res += ToString(RandInt(1000))
	return MD5(res)
}

func RandInt(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}

func ToString(data any) (ret string) {
	switch data.(type) {
	case int64:
		ret = strconv.Itoa(int(data.(int64)))
	case int32:
		ret = strconv.Itoa(int(data.(int32)))
	case int:
		ret = strconv.Itoa(data.(int))
	case float64:
		ret = strconv.FormatFloat(data.(float64), 'e', 10, 64)
	case float32:
		ret = strconv.FormatFloat(data.(float64), 'e', 10, 32)
	}
	return ret
}
