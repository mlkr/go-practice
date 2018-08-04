package problem3

func reverseStr(s string) string {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}

	b := []byte(s)
	i, j := 0, sLen-1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b)
}
