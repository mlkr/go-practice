package problem4

import "strings"

// 最佳
func isSame(str1, str2 string) bool {
	b := make(map[byte]int)

	b1 := []byte(str1)
	b2 := []byte(str2)

	if len(b1) != len(b2) {
		return false
	}

	for _, v := range b1 {
		b[v] += 1
	}

	for _, v := range b2 {
		b[v] -= 1
	}

	for _, v := range b {
		if v > 0 {
			return false
		}
	}

	return true
}

func isSame2(str1, str2 string) bool {
	b := make([]int, 256)

	b1 := []byte(str1)
	b2 := []byte(str2)

	if len(b1) != len(b2) {
		return false
	}

	for _, v := range b1 {
		b[v] += 1
	}

	for _, v := range b2 {
		b[v] -= 1
	}

	for _, v := range b {
		if v > 0 {
			return false
		}
	}

	return true
}

// 来自 https://github.com/lifei6671/interview-go/blob/master/src/q004.go
func isSame3(str1, str2 string) bool {
	sl1 := len([]rune(str1))
	sl2 := len([]rune(str2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}

	for _, v := range str1 {
		if strings.Count(str1, string(v)) != strings.Count(str2, string(v)) {
			return false
		}
	}
	return true
}
