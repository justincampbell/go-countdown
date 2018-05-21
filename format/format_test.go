package format

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	tests := []struct {
		d      time.Duration
		Clock  string
		MinSec string
	}{
		{0, "0:00:00", "0:00"},
		{time.Second, "0:00:01", "0:01"},
		{time.Minute, "0:01:00", "1:00"},
		{time.Hour, "1:00:00", "60:00"},
		{3 * time.Hour, "3:00:00", "180:00"},
		{72 * time.Hour, "72:00:00", "4320:00"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Clock(%s)", tt.d.String()), func(t *testing.T) {
			assert.Equal(t, tt.Clock, Clock(tt.d))
		})

		t.Run(fmt.Sprintf("MinSec(%s)", tt.d.String()), func(t *testing.T) {
			assert.Equal(t, tt.MinSec, MinSec(tt.d))
		})
	}
}
