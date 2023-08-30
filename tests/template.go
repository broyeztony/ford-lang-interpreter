package tests

import (
	interp "ford-lang-interpreter/interp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	input := ``

	actual := interp.Eval(interp.ParseAST(input))
	expected := "<>"

	assert.Equal(t, expected, actual)
}
