
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
		{high: 22.488, low: 21.61, open: 22.3005, close: 21.71516, time: 1123016400, volume: 1878000},
		{high: 22.789, low: 21.61001, open: 21.8, close: 22.59434, time: 1123102800, volume: 2547900},
		{high: 22.699, low: 22.351, open: 22.55, close: 22.64, time: 1123189200, volume: 649700},
		{high: 22.73994, low: 22.45001, open: 22.595, close: 22.56998, time: 1123448400, volume: 587800},
		{high: 22.69, low: 21.8, open: 22.605, close: 22.37001, time: 1123534800, volume: 1431000},
		{high: 22.58, low: 22.26, open: 22.493, close: 22.397, time: 1123621200, volume: 761500},
		{high: 22.442, low: 22.0011, open: 22.397, close: 22.099, time: 1123707600, volume: 1208000},
		{high: 22.16, low: 21.72, open: 22.09, close: 21.937, time: 1123794000, volume: 1130300},