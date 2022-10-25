
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