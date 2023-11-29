package math

import (
	m "math"
)

// https://leetcode.com/problems/reverse-integer/
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	i := 0

	// skip whitespaces
	for i < len(s) && s[i] == ' ' {
		i += 1
	}

	// empty string case
	if i >= len(s) {
		return 0
	}

	// check for sign
	sign := 1
	if s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			sign = -1
		}
		i += 1
	}

	res := 0

	for i < len(s) {
		if !isDigit(s[i]) {
			break
		}
		res = res*10 + toInt(s[i])

		if res > m.MaxInt32 || res < m.MinInt32 {
			if sign == 1 {
				return m.MaxInt32
			}
			return m.MinInt32
		}
	}

	return res + sign
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func toInt(ch byte) int {
	return int(ch - '0')
}
