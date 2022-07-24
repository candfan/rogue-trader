
package indicator

import "math"

const (
	FlatTrend float64 = 0
	UpTrend   float64 = 1
	DownTrend float64 = -1
)

// Trend returns a trend direction.
// It bases on fast (with shorter period) and slow EMA.
// flatMaxDiff allows setting max difference between
// fast and slow EMA when Calculate returns the flat.
type Trend struct {
	fastEMAIndicator Indicator
	slowEMAIndicator Indicator
	flatMaxDiff      float64
}

func NewTrend(
	fastEMAIndicator Indicator,
	slowEMAIndicator Indicator,
	flatMaxDiff float64,
) *Trend {
	return &Trend{
		fastEMAIndicator: fastEMAIndicator,
		slowEMAIndicator: slowEMAIndicator,
		flatMaxDiff:      flatMaxDiff,