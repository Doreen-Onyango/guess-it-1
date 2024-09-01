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
	var diff float64
	for _, num := range data {
		d := num - Mean(data)
		diff += d * d
	}
	return diff / float64(len(data))
}

func StandardDeviation(data []float64) float64 {
	if len(data) == 0 {
		fmt.Println("Provide Data")
		os.Exit(0)
	}
	return math.Sqrt(Variance(data))
}
