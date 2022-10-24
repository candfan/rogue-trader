
package indicator

import (
	"sync"
	"time"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

const oneDay = time.Hour * 24

// VolumeWeightedAveragePrice represents indicator to calculate volume-weighted average price (VWAP).