package utils

import "strconv"

func StringToInt64(s string) int64 {
	r, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return r
}
