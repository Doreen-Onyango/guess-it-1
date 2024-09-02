package calculate

import (
	"math"
	"reflect"
	"testing"
)

func TestMean(t *testing.T) {
	input := []float64{100, 200, 300, 400, 500}
	expected := 300.0
	result := Mean(input)
	roundresult := math.Round(result)
	if !reflect.DeepEqual(roundresult, expected) {
		t.Errorf("test failed. expected : %v, got : %v", expected, roundresult)
	}
}
