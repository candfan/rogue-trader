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

	assert.Nil(t, series.AddCandle(createCandle("2020-06-24