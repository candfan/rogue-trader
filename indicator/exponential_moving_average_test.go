package indicator

import (
	"sync"
	"testing"
	"time"

	"github.com/evsamsonov/trading-timeseries/timeseries"
	"github.com/stretchr/testify/assert"
)

func TestExponentialMovingAverage_Calculate(t *testing.T) {
	series := GetTestSeries()

	tests := []struct {
		name           string
		smoothInterval int
		index          int
		expected       float64
	}{
		{name: "not enough data", smoothInterval: 3, index: 0, expected: 22.104