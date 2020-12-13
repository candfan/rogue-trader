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
		{16, 15, 0.7202