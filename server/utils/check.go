package utils

import (
	"regexp"

	"go.uber.org/zap"
)

const (
	_EMAIL_PATTERN = `\w+@\w+(\.\w+)+`

	_CHINA_TELECOM_PATTERN = `(?:^(?:\+86)?1(?:33|49|53|7[37]|8[019]|9[19])\d{8}$)|(?:^(?:\+86)?1349\d{7}$)|(?:^(?:\+86)?1410\d{7}$)|(?:^(?:\+86)?170[0-2]\d{7}$)`
	_CHINA_UNICOM_PATTERN  = `(?:^(?:\+86)?1(?:3[0-2]|4[56]|5[56]|66|7[156]|8[56])\d{8}$)|(?:^(?:\+86)?170[47-9]\d{7}$)`
	_CHINA_MOBILE_PATTERN  = `(?:^(?:\+86)?1(?:3[4-9]|4[78]|5[0-27-9]|78|8[2-478]|98|95)\d{8}$)|(?:^(?:\+86)?1440\d{7}$)|(?:^(?:\+86)?170[356]\d{7}$)`
)

func CheckEmail(email string) bool {
	b, e := regexp.MatchString(_EMAIL_PATTERN, email)
	if e != nil {
		zap.S().Warn(e)
	}
	return b
}

func CheckPhone(phone string) bool {
	if len(phone) != 11 {
		return false
	}
	return checkPhoneHelper(_CHINA_MOBILE_PATTERN, phone) ||
		checkPhoneHelper(_CHINA_TELECOM_PATTERN, phone) ||
		checkPhoneHelper(_CHINA_UNICOM_PATTERN, phone)
}

func checkPhoneHelper(pattern, phone string) bool {
	b, e := regexp.MatchString(pattern, phone)
	if e != nil {
		zap.S().Warn(e)
	}
	return b
}
