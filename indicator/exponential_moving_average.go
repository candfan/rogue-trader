
package indicator

import (
	"errors"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

// ExponentialMovingAverage represents indicator to calculate Exponential Moving Average (EMA)
type ExponentialMovingAverage struct {
	series         *timeseries.TimeSeries
	smoothInterval int
	maxIndex       int
	cache          []float64
	smooth         float64
}

func NewExponentialMovingAverage(series *timeseries.TimeSeries, smoothInterval int) (*ExponentialMovingAverage, error) {
	if series.Length() == 0 {