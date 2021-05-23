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
		{name: "smoothInterval=3,index=5", period: 3, index