
package indicator

import (
	"math"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

// AverageTrueRange represents indicator to calculate Average True Range (ATR)
// https://en.wikipedia.org/wiki/Average_true_range
type AverageTrueRange struct {
	series       *timeseries.TimeSeries
	period       int
	cache        map[int]float64
	lastComputed int
}

func NewAverageTrueRange(series *timeseries.TimeSeries, period int) Indicator {
	return &AverageTrueRange{series, period, make(map[int]float64), period - 1}
}
