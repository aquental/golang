// Package gigasecond 
// adds a Giga seconds to now
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond add a Giga second to now
func AddGigasecond(t time.Time) time.Time {
	gigaSec := 1000000000
    t = t.Add(time.Second * time.Duration(gigaSec))
	return t
}
