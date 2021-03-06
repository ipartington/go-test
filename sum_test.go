package sum

import "testing"
import "github.com/stretchr/testify/assert"

func TestSum(t *testing.T) {
	total := Add(5, 5)

	assert.Equal(t, total, 10, "they should be equal")
}

func TestSumLarge(t *testing.T) {
	total := Add(500, 501)
	assert.Equal(t, total, 1001, "they should be equal")
}
