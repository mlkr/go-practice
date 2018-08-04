package problem5

import "strings"

func replaceBlank(str string) string {
	b := []byte(str)

	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			temp := append([]byte{'%', '2', '0'}, b[i+1:]...)
			b = append(b[:i], temp...)
		}
	}

	return string(b)
}

func replaceBlank2(str string) string {
	return strings.Replace(str, " ", "%20", -1)
}
