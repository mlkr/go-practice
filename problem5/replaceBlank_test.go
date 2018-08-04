package problem5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	str string
}

type ans struct {
	one string
}

type question struct {
	para
	ans
}

var questions []question

func init() {
	questions = []question{

		question{
			para{"Hello World"},
			ans{"Hello%20World"},
		},
	}
}

func TestReplaceBlank(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(replaceBlank(para.str), ans.one, "输入 %v", para)
	}
}

func TestReplaceBlank2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(replaceBlank2(para.str), ans.one, "输入 %v", para)
	}
}

func BenchmarkReplaceBlank(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceBlank("Hello World")
	}
}

func BenchmarkReplaceBlank2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceBlank2("Hello World")
	}
}
