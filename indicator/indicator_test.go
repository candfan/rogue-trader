
package indicator

import (
	"log"
	"time"

	"github.com/evsamsonov/trading-timeseries/timeseries"
)

const float64EqualityThreshold = 1e-6

type TestCandle struct {
	high   float64
	low    float64
	open   float64
	close  float64
	time   int64
	volume int64
}

func GetTestCandles() []TestCandle {
	return []TestCandle{
		{high: 23, low: 21.27, open: 21.3125, close: 22.1044, time: 1121979600, volume: 4604900},
		{high: 23.31999, low: 22.15, open: 22.15, close: 23.21608, time: 1122238800, volume: 4132600},
		{high: 23.2755, low: 22.111, open: 23.2755, close: 22.20585, time: 1122325200, volume: 2572600},
		{high: 22.58, low: 22.015, open: 22.166, close: 22.2, time: 1122411600, volume: 1441300},
		{high: 22.52, low: 22.13501, open: 22.29969, close: 22.15, time: 1122498000, volume: 1184000},
		{high: 22.3, low: 21.601, open: 22.20443, close: 21.861, time: 1122584400, volume: 1532000},
		{high: 22.1, low: 21.45, open: 21.8499, close: 22.09, time: 1122843600, volume: 1177800},
		{high: 22.38, low: 22.0511, open: 22.111, close: 22.25952, time: 1122930000, volume: 1419400},