package countdown

import "time"

var (
	timeFunc = time.Now
)

// Countdown is a countdown.
type Countdown struct {
	// Interval is the duration in between function invocations.
	Interval time.Duration

	// Duration is how long the countdown should run for.
	Duration time.Duration

	// EndTime is the time the countdown will end. This is set when the
	// countdown is started.
	EndTime time.Time
}

// Func is a function which may be called one or more times by this
// package.
type Func func(*Countdown) error

// For takes a duration and returns a new Countdown.
func For(d, interval time.Duration) *Countdown {
	return &Countdown{
		Interval: interval,
		Duration: d,
	}
}

// Until takes a time and returns a new Countdown.
func Until(t time.Time, interval time.Duration) *Countdown {
	return &Countdown{
		Interval: interval,
		Duration: t.Sub(timeFunc()),
	}
}

// Do counts down for the given duration and calls the CountdownFunc each
// Interval.
func (c *Countdown) Do(f Func) error {
	tick := time.Tick(c.Interval)
	end := time.After(c.Duration)
	c.EndTime = timeFunc().Add(c.Duration)

	err := f(c)
	if err != nil {
		return err
	}

	for {
		select {
		case <-tick:
			err = f(c)
			if err != nil {
				return err
			}
		case <-end:
			return nil
		}
	}
}

// Remaining is the duration of time remaining for this countdown.
func (c *Countdown) Remaining() time.Duration {
	return c.EndTime.Sub(timeFunc())
}
