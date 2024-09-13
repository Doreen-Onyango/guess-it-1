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
	n := len(data)
	if n > 5 {
		n = 5
	}

	// Get the last n elements
	start := len(data) - n
	recentData := data[start:]

	var sum float64
	for _, value := range recentData {
		sum += value
	}
	return sum / float64(len(data))
}
