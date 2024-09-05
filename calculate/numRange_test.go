package calculate

import (
	"math"
	"reflect"
	"testing"
)

func TestCalculateRange(t *testing.T) {
	input := []float64{100, 200, 300, 400, 500}
	expected1 := 17.0
	expected2 := 583.0
	res1, res2 := CalculateRange(input)
	roundres1 := math.Round(res1)
	roundres2 := math.Round(res2)
	if !reflect.DeepEqual(roundres1, expected1) {
		t.Errorf("test failed. expected : %v got : %v", expected1, roundres1)
	}
	if !reflect.DeepEqual(roundres2, expected2) {
		t.Errorf("test failed. expected : %v got : %v", expected2, roundres2)
	}
}
