package model

import (
	"github.com/jthomperoo/holtwinters"
)

type ExpSmooth struct {
	alpha float64
	beta  float64
	gamma float64
}

func NewExpSmooth(alpha, beta, gamma float64) *ExpSmooth {
	return &ExpSmooth{
		alpha: alpha,
		beta:  beta,
		gamma: gamma,
	}
}

func (es *ExpSmooth) PredictLTV(y []float64, day int) float64 {
	res, _ := holtwinters.PredictAdditive(y, len(y), 0.1, 0.25, 0.35, 60-len(y))
	return res[len(res)-1]
}
