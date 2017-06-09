package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

// AddGigasecond add 10^9 seconds to a given Time struct and return it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e18)
}
