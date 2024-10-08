package calculate

func CalculateRange(num []float64) (float64, float64) {
	if len(num) == 0 {
		return 0, 0
	}

	meanDiff := Mean(num)
	stdDiff := StandardDeviation(num)

	// Capture variability around the predicted value.
	minLimit := meanDiff - 2*stdDiff
	maxLimit := meanDiff + 2*stdDiff

	return minLimit, maxLimit
}
