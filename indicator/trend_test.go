package indicator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrend_Calculate(t *testing.T) {
	tests := []st