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
		