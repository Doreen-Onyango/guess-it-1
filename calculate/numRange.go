package calculate

func CalculateRange(num []float64) (float64, float64) {
	if len(num) < 2 {
		return 0, 1000
	}

	dif := make([]float64, len(num)-1)
	for i := 1; i < len(num); i++ {
		dif[i-1] = num[i] - num[i-1]
	}

	meanDiff := Mean(dif)
	stdDiff := StandardDeviation(dif)

	lastNum := num[len(num)-1]
	predict := lastNum + meanDiff

	minLimit := predict - 3*stdDiff
	maxLimit := predict + 3*stdDiff

	return minLimit, maxLimit
}
