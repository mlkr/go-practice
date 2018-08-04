package problem6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans Coor

type question struct {
	para
	ans
}

func TestGetCoordinate(t *testing.T) {
	questions := []question{

		question{
			para{"R2(LF)"},
			ans{-1, 1},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		res := getCoordinate(para.one)

		if res.x != ans.x || res.y != ans.y {
			t.Fatalf("输入 %v", para)
		}
	}
}

func TestAnalyze(t *testing.T) {
	type para struct {
		one string
	}

	type ans struct {
		one string
	}

	type question struct {
		para
		ans
	}

	ast := assert.New(t)

	questions := []question{

		question{
			para{"R2(LF)"},
			ans{"RLFLF"},
		},

		question{
			para{"R2(B2(LF2(RF)))F"},
			ans{"RBLFRFRFLFRFRFBLFRFRFLFRFRFF"},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(analyze(para.one), []byte(ans.one), "输入 %v", para)
	}
}
