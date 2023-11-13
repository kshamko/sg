package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLEPredictLTV(t *testing.T) {

	testCases := []struct {
		name     string
		ltvs     []float64
		day      int
		expected float64
	}{
		{
			name:     "2 points",
			ltvs:     []float64{0, 1},
			day:      3,
			expected: 2,
		},
		{
			name:     "3 points",
			ltvs:     []float64{0, 1, 3},
			day:      4,
			expected: 4.333333333333333,
		},
		{
			name:     "empty",
			ltvs:     []float64{},
			day:      3,
			expected: 0,
		},
	}

	le := NewLinExtrapol()

	for _, testCase := range testCases {
		predicted := le.PredictLTV(testCase.ltvs, testCase.day)
		assert.Equal(t, testCase.expected, predicted)
	}

}
