package indicator

import (
	"math"
	"testing"
)

func TestAverageTrueRange_Calculate(t *testing.T) {
	series := GetTestSeries()

	testCases := []struct {
		