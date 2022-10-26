
package indicator

import (
	"sync"
	"time"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

const oneDay = time.Hour * 24

// VolumeWeightedAveragePrice represents indicator to calculate volume-weighted average price (VWAP).
// More details https://en.wikipedia.org/wiki/Volume-weighted_average_price
type VolumeWeightedAveragePrice struct {
	series *timeseries.TimeSeries
	cache  *vwapCache
}

// NewVolumeWeightedAveragePrice creates VolumeWeightedAveragePrice
func NewVolumeWeightedAveragePrice(series *timeseries.TimeSeries) *VolumeWeightedAveragePrice {
	return &VolumeWeightedAveragePrice{
		series: series,
		cache:  newVwapCache(),
	}
}

// Calculate returns VWAP value for candle with given index
func (v *VolumeWeightedAveragePrice) Calculate(index int) float64 {
	unit, ok := v.cache.get(index)
	if ok {
		return unit.vwap
	}

	day := v.series.Candle(index).Time.Truncate(oneDay)

	startIndex, unit := v.findLastCalculated(index, day)
	volumeTotal := unit.volumeTotal