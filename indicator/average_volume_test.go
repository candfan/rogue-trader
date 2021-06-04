package indicator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageVolume_Calculate(t *testing.T) {
	series := GetTestSeries()

	tests := []struct {
		name     string
		period   int
		index    int
		expected float64
	}{
		{name: "not enough data", period: 3, index: 1, expected: 0},
		{name: "smoothInterval=3,index=5", period: 3, index: 5, expected: 1385766.66},
		{name: "smoothInterval=14,index=13", period: 14, index: 13, expected: 1851464.2857},
	}

	for _, test := range tests {
