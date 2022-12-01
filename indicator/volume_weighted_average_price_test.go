package indicator

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/evsamsonov/trading-timeseries/timeseries"
	"github.com/stretchr/testify/assert"
)

func TestVolumeWeightedAveragePrice_Calculate(t *testing.T) {
	series := timeseries.New()

	assert.Nil(t, series.AddCandle(createCandle("2020-06-24T00:00:00+00:00", 1, 2, 3, 100)))
	assert.Nil(t, series.AddCandle(createCandle("2020-06-25T00:00:00+00:00", 4, 5, 6, 