package gigasecond

import "time"

const GIGASECOND = 1000000000

// Given time.Time
// Return time.Time a GIGASECOND later
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * GIGASECOND)
}
