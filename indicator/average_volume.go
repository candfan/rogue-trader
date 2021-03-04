package indicator

import "github.com/evsamsonov/trading-timeseries/timeseries"

// AverageVolume is a indicator to calculate average volume
// for given smoothInterval using simple moving average
type AverageVolume struct {
	series *timeseries.TimeSeries
	period int
	cache  map[int]float64
}

func NewAverageVolume(series *timeseries.TimeSeries, period int) *AverageVolume {
	return &AverageVolume{
		series: series,
		pe