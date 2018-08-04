package problem3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestReverseStr(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{"mlkr"},
			ans{"rklm"},
		},

		question{
			para{"mlkr yun"},
			ans{"nuy rklm"},
		},

		question{
			para{"This is golang"},
			ans{"gnalog si sihT"},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(reverseStr(para.one), ans.one, "输入 %v", para)
	}
}
