package calculate

import (
	"math"
	"reflect"
	"testing"
)

func TestCalculateRange(t *testing.T) {
	input := []float64{100, 200, 300, 400, 500}
	expected := 600.0
	res1, res2 := CalculateRange(input)
	roundres1 := math.Round(res1)
	roundres2 := math.Round(res2)
	if !reflect.DeepEqual(roundres1, expected) {
		t.Errorf("test failed. expected : %v, got : %v", expected, roundres1)
	}
	if !reflect.DeepEqual(roundres2, expected) {
		t.Errorf("test failed. expected : %v, got : %v", expected, roundres2)
	}
}
