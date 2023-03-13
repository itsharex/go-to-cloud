package utils

import (
	"regexp"
	"strings"
)

func isMatch(pattern, o string) bool {
	match, err := regexp.MatchString(pattern, o)
	if err != nil {
		return false
	}
	return match
}

func IsValidEmail(email string) bool {
	return isMatch(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, strings.ToLower(email))
}

func IsValidMobile(mobile string) bool {
	return isMatch(`^1[3-9]\d{9}$`, mobile)
}
