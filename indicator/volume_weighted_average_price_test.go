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
	series := timeseries.New(