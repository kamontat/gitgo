package utils

import "strings"

func updateSelectResult(str string) string {
	arr := strings.Split(str, ":")
	if len(arr) > 0 {
		s := arr[0]
		return strings.TrimSpace(s)
	}

	return str
}
