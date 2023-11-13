package predict

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultAgregator(t *testing.T) {

	testCases := []struct {
		ltvs     []float64
		counts   []uint
		expected float64
	}{
		{
			ltvs:     []float64{1, 20},
			counts:   []uint{1, 10},
			expected: float64(21) / float64(11),
		},
	}

	for _, testCase := range testCases {
		agg := NewDefaultAggregator()

		for i, ltv := range testCase.ltvs {
			agg.Collect("test", ltv, testCase.counts[i])
		}

		res := agg.Result()

		assert.Equal(t, res[0].Value, testCase.expected)

	}
}
