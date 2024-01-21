package clock

import "fmt"

const (
    hour_per_day int = 24
    min_per_hour int = 60
    sec_per_min int = 60
)

// Define the Clock type here.
type Clock struct {
    hour int
    min int
}

func norm(h int, m int) (int, int) {
	for m >= min_per_hour {
		m -= min_per_hour
		h++
	}
	for h >= hour_per_day {
		h -= hour_per_day
	}
	for m < 0 {
		m += min_per_hour
		h--
	}
	for h < 0 {
		h += hour_per_day
	}
	return h, m
}

func New(h, m int) Clock {
    h, m = norm(h, m)
    return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.min + m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.min - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

