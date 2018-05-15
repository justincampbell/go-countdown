package countdown

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	now time.Time
)

func init() {
	now = time.Time{}
	timeFunc = func() time.Time { return now }
}

func TestFor(t *testing.T) {
	c := For(time.Minute, time.Second)
	assert.Equal(t, time.Second, c.Interval)
	assert.Equal(t, time.Minute, c.Duration)
}

func TestUntil(t *testing.T) {
	c := Until(now.Add(time.Nanosecond), time.Second)
	assert.Equal(t, time.Second, c.Interval)
	assert.Equal(t, time.Nanosecond, c.Duration)
}

func TestCountdown_Do(t *testing.T) {
	duration := 3 * time.Millisecond
	interval := time.Millisecond
	c := For(duration, interval)

	t.Run("EndTime", func(t *testing.T) {
		err := c.Do(func(c *Countdown) error { return nil })

		assert.Nil(t, err)
		assert.Equal(t, timeFunc().Add(duration), c.EndTime,
			"sets the end time")
	})

	t.Run("Func", func(t *testing.T) {
		called := 0
		err := c.Do(func(c *Countdown) error {
			called++
			return nil
		})

		assert.Nil(t, err)
		assert.Equal(t, 4, called,
			"calls function for each interval plus 1 at the start")
	})

	t.Run("error", func(t *testing.T) {
		called := 0
		err := c.Do(func(c *Countdown) error {
			called++
			return fmt.Errorf("Error")
		})

		assert.NotNil(t, err,
			"returns the error")
		assert.Equal(t, 1, called,
			"stops after the error")
	})
}

func TestCountdown_Remaining(t *testing.T) {
	c := For(3*time.Millisecond, time.Millisecond)

	assert.Equal(t, 0, c.Remaining(),
		"returns zero when not running")

	c.Do(func(c *Countdown) error {
		assert.Equal(t, c.EndTime.Sub(timeFunc()), c.Remaining(),
			"returns the time remaining")

		return fmt.Errorf("Error")
	})
}
