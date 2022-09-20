package indicator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrend_Calculate(t *testing.T) {
	tests := []struct {
		name    string
		fastVal float64
		slowVal float64
		want    float64
	}{
		{
			name:    "up trend",
			fastVal: 1.61,
			slowVal: 1.0,
			want:    UpTrend,
		},
		{
			name:    "down trend",
			fastVal: 1.0,
			slowVal: 1.61,
			want:    DownTrend,
		},
		{
			name:    "flat trend",
			fastVal: 1.0,
			slowVal: 1.6,
			want:    FlatTrend,
		},
		{
			name:    "flat trend 2",
			fastVal: 1.6,
			slowVal: 1.0,
			want:    FlatTrend,
		},
	}

	for _, tt := range tests 