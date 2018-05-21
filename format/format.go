package format

import (
	"fmt"
	"time"
)

// Clock returns the duration formatted as h:mm:ss.
func Clock(d time.Duration) string {
	if d < 0 {
		return "0:00"
	}

	s := int(d.Seconds() + 0.5)

	m := s / 60
	s = s % 60

	h := m / 60
	m = m % 60

	return fmt.Sprintf("%d:%02d:%02d", h, m, s)
}

// MinSec returns the duration formatted as mm:ss.
func MinSec(d time.Duration) string {
	if d < 0 {
		return "0:00"
	}

	s := int(d.Seconds() + 0.5)

	m := s / 60
	s = s % 60

	return fmt.Sprintf("%d:%02d", m, s)
}
