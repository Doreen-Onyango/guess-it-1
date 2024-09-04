package calculate

import (
	"fmt"
	"math"
	"os"
)

func Variance(data []float64) float64 {
	if len(data) == 0 {
		fmt.Println("Provide Data")
		os.Exit(0)
	}
	mean := Mean(data)
	out := []float64{}

	for _, num := range data {
		dev := num - mean
		out = append(out, dev*dev)
	}
	return Mean(out)
}

func StandardDeviation(data []float64) float64 {
	if len(data) == 0 {
		fmt.Println("Provide Data")
		os.Exit(0)
	}
	return math.Sqrt(Variance(data))
}
