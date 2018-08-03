package problem1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintNumLetter(t *testing.T) {
	ast := assert.New(t)

	ans := "12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728"

	ast.Equal(printNumLetter(), ans)
}
