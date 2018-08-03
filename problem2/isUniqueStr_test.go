package problem2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one bool
}

type question struct {
	para
	ans
}

func TestIsUniqueStr(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{"abfgasdgfdgfdfdfds"},
			ans{false},
		},

		question{
			para{"abcdefG"},
			ans{true},
		},

		question{
			para{""},
			ans{false},
		},

		question{
			para{"a"},
			ans{true},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		// ast.Equal(isUniqueStr(para.one), ans.one, "输入 %v", para)
		ast.Equal(isUniqueStr5(para.one), ans.one, "输入 %v", para)
	}
}

func BenchmarkIsUniqueStr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isUniqueStr("abfgasdgfdgfdfdfds")
	}
}

func BenchmarkIsUniqueStr2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isUniqueStr2("abfgasdgfdgfdfdfds")
	}
}

func BenchmarkIsUniqueStr3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isUniqueStr3("abfgasdgfdgfdfdfds")
	}
}

func BenchmarkIsUniqueStr4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isUniqueStr4("abfgasdgfdgfdfdfds")
	}
}

func BenchmarkIsUniqueStr5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isUniqueStr4("abfgasdgfdgfdfdfds")
	}
}
