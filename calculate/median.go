package calculate

import (
	"fmt"
	"os"
	"sort"
)

// Finding median by sorting the data and taking the middle if the
// length of data is odd and dividing the middle two values if the
// length of data is even.
func Median(data []float64) float64 {
	if len(data) == 0 {
		fmt.Println("Empty data set provided")
		os.Exit(0)
	}
	sort.Float64s(data)
	mid := len(data) / 2
	if len(data)%2 != 0 {
		return data[mid]
	} else {
		return (data[mid-1] + data[mid]) / 2
	}
}
