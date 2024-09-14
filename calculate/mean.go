package calculate

import (
	"fmt"
	"os"
)

// Computing average by adding the whole data and dividing with the length of that data.
func Mean(data []float64) float64 {
	if len(data) == 0 {
		fmt.Println("Provide Data")
		os.Exit(0)
	}
	var sum float64
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}
