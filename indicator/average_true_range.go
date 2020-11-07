
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

func (ind *AverageTrueRange) Calculate(index int) float64 {
	if index < ind.period-1 {
		return 0
	}

	if val, ok := ind.cache[index]; ok {
		return val
	}

	for i := ind.lastComputed; i <= index; i++ {
		ind.cache[i] = ind.doCalculate(i)