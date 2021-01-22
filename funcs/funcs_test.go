package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// struct testCase{
// 	s string,
// 	op string,
// 	a int,
// 	b int,
// 	res int
// }

// tests = make([]testCase,0,10);

func TestAdd(t *testing.T) {

	assert.Equal(t, 10, add(7, 3))
	assert.Equal(t, 23, add(13, 10))
}

func TestSubtract(t *testing.T) {
	assert.Equal(t, 4, subtract(7, 3))
	assert.Equal(t, 3, subtract(13, 10))
	assert.Equal(t, -2, subtract(1, 3))
	assert.Equal(t, 0, subtract(13, 13))
	assert.Equal(t, -13, subtract(0, 13))
	assert.Equal(t, 13, subtract(13, 0))
}

func TestProduct(t *testing.T) {
	assert.Equal(t, 13, product(1, 13))
	assert.Equal(t, 0, product(13, 0))
	assert.Equal(t, -13, product(-13, 1))
	assert.Equal(t, 26, product(13, 2))
}

func TestDivision(t *testing.T) {
	assert.Equal(t, 0, division(1, 13))
	assert.Equal(t, 13, division(13, 1))
	assert.Equal(t, -13, division(-13, 1))
	assert.Equal(t, 6, division(13, 2))
}

func TestModulo(t *testing.T) {
	assert.Equal(t, 1, modulo(1, 13))
	assert.Equal(t, 0, modulo(-13, 1))
	assert.Equal(t, 1, modulo(13, 2))
}
