package indicator

import (
	"math"
	"testing"
)

func TestAverageTrueRange_Calculate(t *testing.T) {
	series := GetTestSeries()

	testCases := []struct {
		period   int
		index    int
		expected float64
	}{
		{16, 15, 0.720238},
		{16, 14, 0.0},
		{8, 15, 0.650576},
	}

	for _, test := range testCases {
		atrIndicator := NewAverageTrueRange(series, test.period)
		atr := atrIndicator.Calculate(test.index)
		expectedAtr :=