package math

import (
	m "math"
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	// 123

	// res := 0, x = 123

	// x = 123; res = (0 * 10) + (3) => 3

	// x = 12; res = (3 * 10) + (2) => 32

	// x = 1; res = (32 * 10) + (1) => 321

	res := 0

	for x > 0 {
		res = (res * 10) + (x % 10)
		if res > m.MaxInt32 {
			return 0
		}
		x = x / 10
	}

	return res * sign
}
