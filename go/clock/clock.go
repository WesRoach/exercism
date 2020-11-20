package clock

import "fmt"

const (
	minutesInDay = 1440
)

// Clock holds time in minutes
type Clock struct {
	minutes int
}

// New creates a new Clock given hour & minutes
func New(hour int, minutes int) Clock {
	return Clock{0}.Add(hour*60 + minutes)
}

// String's Clock struct
func (clock Clock) String() string {
	return fmt.Sprintf("%02v:%02v", clock.minutes/60, clock.minutes%60)
}

// Add adds minutes to an existing Clock
func (clock Clock) Add(minutes int) Clock {
	minutes = (clock.minutes + (minutes % minutesInDay)) % minutesInDay
	if minutes < 0 {
		minutes += minutesInDay
	}
	return Clock{minutes}
}

// Subtract substracts minutes from an existing Clock
func (clock Clock) Subtract(minutes int) Clock {
	return clock.Add(-minutes)
}
