package calculate

import (
	"math"
	"reflect"
	"testing"
)

func TestVariance(t *testing.T) {
	input := []float64{100, 200, 300, 400, 500}
	expected := 20000.0
	result := Variance(input)
	roundresult := math.Round(result)
	if !reflect.DeepEqual(roundresult, expected) {
		t.Errorf("test failed. expected : %v, got : %v", expected, roundresult)
	}
}

func TestStandardDeviation(t *testing.T) {
	input := []float64{100, 200, 300, 400, 500}
	expected := 141.0
	result := StandardDeviation(input)
	roundresult := math.Round(result)
	if !reflect.DeepEqual(roundresult, expected) {
		t.Errorf("test failed. expected : %v, got : %v", expected, roundresult)
	}
}
