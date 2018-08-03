package problem2

import (
	"math/rand"
	"strings"
)

func isUniqueStr(str string) bool {
	strLen := len(str)
	if strLen == 0 {
		return false
	}

	strB := []byte(str)

	return quickSort(strB)
}

// 使用修改后的快排检查字符串中有没有重复字母
func quickSort(b []byte) bool {
	bLen := len(b)
	if bLen <= 1 {
		return true
	}

	point := rand.Intn(bLen)
	return part(b, point)
}

func part(b []byte, point int) bool {
	bLen := len(b)
	pointNum := b[point]
	b[0], b[point] = b[point], b[0]

	i, j := 1, bLen-1

	for i < j {

		if b[i] == pointNum {
			return false
		}

		for b[i] < pointNum && i < bLen-1 {
			i++
		}

		for pointNum < b[j] && j > 0 {
			j--
		}

		if i < j {
			b[i], b[j] = b[j], b[i]
		}
	}

	b[0], b[j] = b[j], b[0]

	res1 := quickSort(b[:j])
	if res1 == false {
		return false
	}

	res2 := quickSort(b[j+1:])

	return res2
}

// 简化
func isUniqueStr2(str string) bool {
	strLen := len(str)
	if strLen == 0 {
		return false
	}

	strB := []byte(str)

	return isUnique(strB)
}

func isUnique(b []byte) bool {
	bLen := len(b)
	if bLen <= 1 {
		return true
	}

	point := rand.Intn(bLen)
	pointNum := b[point]

	for i := 0; i < point; i++ {
		if b[i] == pointNum {
			return false
		}
	}

	for j := bLen - 1; j > point; j-- {
		if b[j] == pointNum {
			return false
		}
	}

	res1 := isUnique(b[:point])
	if res1 == false {
		return false
	}

	res2 := isUnique(b[point+1:])

	return res2
}

func isUniqueStr3(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != strings.LastIndex(s, string(v)) {
			return false
		}
	}
	return true
}

func isUniqueStr4(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

// 最佳
func isUniqueStr5(s string) bool {
	sLen := len(s)
	if sLen > 256 || sLen == 0 {
		return false
	}

	for i := 0; i < sLen; i++ {

		j := sLen - 1
		for s[j] != s[i] && j > 0 {
			j--
		}

		if i != j {
			return false
		}

		i++
	}

	return true
}
