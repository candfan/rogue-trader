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
		{name: "n