/*
Package gigasecond contains AddGigasecond function? which add 10^9 seconds to time parameter
*/
package gigasecond

import "time"

// Gigasecond is 10^9 seconds value
const Gigasecond = 1e9 * time.Second

// AddGigasecond add 10^9 seconds to time parameter
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(Gigasecond)
	return t
}
