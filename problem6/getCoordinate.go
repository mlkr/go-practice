package problem6

import (
	"strconv"
)

type Coor struct {
	x int
	y int
}

// 输入命令, 得机器人的坐标
func getCoordinate(cmd string) Coor {
	b := analyze(cmd)
	bLen := len(b)

	lCount, rCount := 0, 0
	fCount, bCount := 0, 0
	coor := &Coor{0, 0}

	for i := 0; i < bLen; i++ {

		switch {
		case b[i] == 'L' || b[i] == 'R':
			if b[i] == 'L' {
				lCount++
			} else {
				rCount++
			}
		default:
			for {
				if b[i] == 'F' {
					fCount++
				} else {
					bCount++
				}

				if i+1 > bLen-1 || b[i+1] == 'L' || b[i+1] == 'R' {
					break
				}

				i++

			}

			move := fCount - bCount

			// 方向
			// 		0
			// 3		1
			// 		2
			dire := (rCount - lCount) % 4
			if dire < 0 {
				dire += 4
			}

			coor.CoorCount(move, dire)
			fCount, bCount = 0, 0
		}
	}

	return *coor
}

// 计算坐标
func (coor *Coor) CoorCount(move, dire int) {
	if move < 0 {
		move = move - (move * 2)
		dire = (dire + 2) % 4
	}

	switch dire {
	case 0:
		// 上
		for i := 0; i < move; i++ {
			coor.y++
		}
	case 1:
		// 右
		for i := 0; i < move; i++ {
			coor.x++
		}
	case 2:
		// 下
		for i := 0; i < move; i++ {
			coor.y--
		}
	case 3:
		// 左
		for i := 0; i < move; i++ {
			coor.x--
		}
	}
}

// 解析命令
func analyze(cmd string) []byte {
	b := []byte(cmd)

	nums := []int{}
	leftBracket := []int{}

	for i := 0; i < len(b); i++ {

		switch {
		case '0' <= b[i] && b[i] <= '9':
			// 处理倍数
			start := i
			for b[i+1] != '(' {
				i++
			}
			end := i

			num, _ := strconv.Atoi(string(b[start : end+1]))
			nums = append(nums, num)

			b = append(b[:start], b[end+1:]...)

			i = start - 1
		case b[i] == '(':
			// 处理 '('
			leftBracket = append(leftBracket, i)
		case b[i] == ')':
			// 处理 ')'
			leftLimit := leftBracket[len(leftBracket)-1]
			leftBracket = leftBracket[:len(leftBracket)-1]

			subStr := b[leftLimit+1 : i]
			times := nums[len(nums)-1]
			nums = nums[:len(nums)-1]

			repStr := repeatStr(subStr, times)

			temp := append(repStr, b[i+1:]...)
			b = append(b[:leftLimit], temp...)

			i = i - 2 + len(repStr) - len(subStr)
		}
	}

	return b
}

// 重复字符串
func repeatStr(str []byte, times int) []byte {
	res := []byte{}

	for i := 0; i < times; i++ {
		res = append(res, str...)
	}

	return res
}
