package calculate

import (
	"math"
	"reflect"
	"testing"
)

func TestMedian(t *testing.T) {
	input := []float64{100, 200, 300, 400, 500}
	expected := 300.0
	result := Median(input)
	roundresult := math.Round(result)
	if !reflect.DeepEqual(roundresult, expected) {
		t.Errorf("test failed. expected : %v, got : %v", expected, roundresult)
	}
}
