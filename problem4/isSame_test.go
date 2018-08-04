package problem4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	str1 string
	str2 string
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

var questions []question

func init() {
	questions = []question{

		question{
			para{"This is golang", "gnalog si sihT"},
			ans{true},
		},

		question{
			para{"Here you are", "Are you here"},
			ans{false},
		},

		question{
			para{"This is golang1.1", "This is golang1"},
			ans{false},
		},
	}
}

func TestIsSame(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(isSame(para.str1, para.str2), ans.one, "输入 %v", para)
	}
}

func TestIsSame2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(isSame2(para.str1, para.str2), ans.one, "输入 %v", para)
	}
}

func TestIsSame3(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(isSame3(para.str1, para.str2), ans.one, "输入 %v", para)
	}
}

func BenchmarkIsSame(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSame("This is golang1.1", "This is golang1")
	}
}

func BenchmarkIsSame2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSame2("This is golang1.1", "This is golang1")
	}
}

func BenchmarkIsSame3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSame3("This is golang1.1", "This is golang1")
	}
}
