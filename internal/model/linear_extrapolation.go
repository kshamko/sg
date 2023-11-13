package model

type LinExtrapol struct{}

func NewLinExtrapol() *LinExtrapol {
	return &LinExtrapol{}
}

func (le *LinExtrapol) PredictLTV(y []float64, day int) float64 {

	if len(y) == 0 {
		return 0
	}

	x := make([]float64, len(y))
	for i := 0; i < len(y); i++ {
		x[i] = float64(i)
	}

	n := len(x)

	// Calculate the sums needed for slope and y-intercept
	sumX, sumY, sumXY, sumXSquare := 0.0, 0.0, 0.0, 0.0
	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXSquare += x[i] * x[i]
	}

	// Calculate slope (m) and y-intercept (b)
	m := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumXSquare - sumX*sumX)
	b := (sumY - m*sumX) / float64(n)

	// Perform extrapolation for xNew
	yNew := m*float64(day-1) + b
	return yNew

}
