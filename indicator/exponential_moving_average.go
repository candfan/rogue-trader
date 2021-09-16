
package indicator

import (
	"errors"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

// ExponentialMovingAverage represents indicator to calculate Exponential Moving Average (EMA)
type ExponentialMovingAverage struct {
	series         *timeseries.TimeSeries